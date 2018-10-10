package redis

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var (
	configs = make(map[string]*Config)
	pools   = make(map[string]*redis.Pool)
)

func init() {
	var redisConfigs map[string]Config
	err := viper.UnmarshalKey(fmt.Sprintf("databases.redis"), &redisConfigs)
	if err != nil {
		panic(err)
	}
	for db, config := range redisConfigs {
		configs[db] = &config
	}
}

type Config struct {
	Host     string      `mapstructure:"host"`
	Port     uint16      `mapstructure:"port"`
	Password string      `mapstructure:"password"`
	Database uint        `mapstructure:"database"`
	Pool     *PoolConfig `mapstructure:"pool"`
}

type PoolConfig struct {
	MaxIdle     int           `mapstructure:"max_idle"`
	MaxActive   int           `mapstructure:"max_active"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

func (c Config) String() string {
	var addr bytes.Buffer
	addr.WriteString(c.Host)
	addr.WriteString(":")
	addr.WriteString(strconv.Itoa(int(c.Port)))
	return addr.String()
}

type Client struct {
	Id   string
	conn redis.Conn
}

func GetConfig(conn string) (*Config, error) {
	if cfg, ok := configs[conn]; ok {
		return cfg, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid connection: %s", conn))
}

func NewClient(db string) *Client {
	cfg, err := GetConfig(db)
	if err != nil {
		panic(err)
	}

	if cfg.Pool != nil {
		if p, ok := pools[db]; !ok {
			p = &redis.Pool{
				MaxIdle:     cfg.Pool.MaxIdle,
				MaxActive:   cfg.Pool.MaxActive,
				IdleTimeout: cfg.Pool.IdleTimeout * time.Second,
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", cfg.String(), redis.DialDatabase(int(cfg.Database)))
				},
			}
			pools[db] = p
		}

		return &Client{Id: "", conn: pools[db].Get()}
	} else {
		conn, err := redis.Dial("tcp", cfg.String(), redis.DialDatabase(int(cfg.Database)))
		if err != nil {
			panic(err)
		}

		if cfg.Password != "" {
			_, err := conn.Do("AUTH", cfg.Password)
			if err != nil {
				panic(err)
			}
		}
		return &Client{Id: db, conn: conn}
	}

}

func (c Client) Select(db uint64) error {
	_, err := c.conn.Do("SELECT", db)
	return err
}

func (c Client) GetDatabases() (map[uint64]string, error) {
	var databases = make(map[uint64]string)

	reply, err := c.conn.Do("INFO", "Keyspace")
	keyspace, err := redis.String(reply, err)
	keyspace = strings.Trim(keyspace[12:], "\n")
	keyspaces := strings.Split(keyspace, "\r")

	for _, db := range keyspaces {
		strs := strings.Split(db, ":")
		strs[0] = strings.Trim(strs[0], "\n")
		if strs[0] == "" {
			continue
		}

		dbi, _ := strconv.ParseUint(strs[0][2:], 10, 64)
		databases[dbi] = strs[1]
	}
	return databases, err
}

func (c Client) Scan(cursor *uint64, match string, limit uint64) ([]string, error) {
	reply, err := c.conn.Do("SCAN", *cursor, "MATCH", match, "COUNT", limit)
	result, err := redis.Values(reply, err)

	var keys []string

	for _, v := range result {
		switch v.(type) {
		case []uint8:
			*cursor, _ = redis.Uint64(v, nil)
		case []interface{}:
			keys, _ = redis.Strings(v, nil)
		}
	}
	return keys, err
}

func (c Client) Ttl(key string) (int64, error) {
	reply, err := c.conn.Do("TTL", key)
	ttl, err := redis.Int64(reply, err)
	return ttl, err
}

func (c Client) SerializedLength(key string) (uint64, error) {
	reply, err := c.conn.Do("DEBUG", "OBJECT", key)
	debug, err := redis.String(reply, err)

	if err != nil {
		return 0, err
	}

	debugs := strings.Split(debug, " ")
	items := strings.Split(debugs[4], ":")

	return strconv.ParseUint(items[1], 10, 64)
}

func (c Client) HmSet(key string, kv map[string]interface{}) error {
	var args []interface{}
	args = append(args, key)
	for k, v := range kv {
		args = append(args, k, v)
	}
	_, err := c.conn.Do("HMSET", args...)
	return err
}

func (c Client) HSet(key string, field string, value interface{}) error {
	_, err := c.conn.Do("HSET", key, field, value)
	return err
}

func (c Client) Set(key string, value interface{}) error {
	_, err := c.conn.Do("SET", key, value)
	return err
}

func (c Client) Setex(key string, expire int, value interface{}) error {
	_, err := c.conn.Do("SETEX", key, expire, value)
	return err
}

func (c Client) HGetAll(key string, value interface{}) error {
	reply, err := c.conn.Do("HGETALL", key)
	v, err := redis.Values(reply, err)
	err = redis.ScanStruct(v, value)
	return err
}

func (c Client) Get(key string) (string, error) {
	val, err := redis.String(c.conn.Do("GET", key))
	return val, err
}

func (c Client) GetInfo(section string) (map[string]string, error) {
	section = strings.Title(strings.ToLower(section))
	info := make(map[string]string)
	reply, err := c.conn.Do("INFO", section)
	if err != nil {
		return nil, err
	}
	strReply := string(reply.([]byte))
	strReply = strings.Replace(strReply, "\r\n", "**", -1)
	expr := fmt.Sprintf(".* %s(.*)", section)
	reg, _ := regexp.Compile(expr)
	subStr := reg.FindStringSubmatch(strReply)[1]
	subStr = strings.Trim(subStr, "*")
	sliceSubstr := strings.Split(subStr, "**")

	for _, s := range sliceSubstr {
		sslice := strings.Split(s, ":")
		info[sslice[0]] = sslice[1]
	}
	return info, nil
}

type PipeLiningCmd struct {
	Cmd  string
	Args interface{}
}

type PipeLiningResult struct {
	Reply interface{}
	Error error
}

func (c Client) Pipelining(cmds []PipeLiningCmd) []PipeLiningResult {
	if len(cmds) == 0 {
		panic(errors.New("no cmds for pipelining"))
	}
	for _, cmd := range cmds {
		switch cmd.Args.(type) {
		case []interface{}:
			err := c.conn.Send(cmd.Cmd, cmd.Args.([]interface{})...)
			if err != nil {
				return []PipeLiningResult{}
			}
		default:
			err := c.conn.Send(cmd.Cmd, cmd.Args)
			if err != nil {
				return []PipeLiningResult{}
			}
		}
	}
	c.conn.Flush()

	vals := make([]PipeLiningResult, len(cmds))
	for id, _ := range cmds {
		r := PipeLiningResult{}
		r.Reply, r.Error = c.conn.Receive()
		vals[id] = r
	}
	return vals
}

// 自己创建的客户端一定要自己使用关闭
func (c Client) Close() error {
	return c.conn.Close()
}
