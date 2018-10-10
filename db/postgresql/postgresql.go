package postgresql

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"gitlab.kucoin.net/golang/framework/config"
)

var (
	configs map[string]*Config
	dbs     = make(map[string]*gorm.DB)
)

type Config struct {
	Host            string
	Port            uint
	Database        string
	Username        string
	Password        string
	SearchPath      string `mapstructure:"search_path"`
	SslMode         string `mapstructure:"sslmode"`
	SslCert         string `mapstructure:"sslcert"`
	SslKey          string `mapstructure:"sslkey"`
	SslRootCert     string `mapstructure:"sslrootcert"`
	ConnectTimeout  uint   `mapstructure:"connect_timeout"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
}

func init() {
	err := viper.UnmarshalKey("databases.postgresql", &configs)
	if err != nil {
		panic(err)
	}
	for conn, cfg := range configs {
		if cfg.Username == "" {
			cfg.Username = "root"
		}
		if cfg.Host == "" {
			cfg.Host = "127.0.0.1"
		}
		if cfg.Port == 0 {
			cfg.Port = 5432
		}
		if cfg.SearchPath == "" {
			cfg.SearchPath = "\"$user\",\\ public"
		}
		if cfg.SslMode == "" {
			cfg.SslMode = "disable"
		}
		if _, err := Connection(conn); err != nil {
			panic(err)
		}
	}
}

func Connection(conn string) (*gorm.DB, error) {
	if db, ok := dbs[conn]; ok {
		return db, nil
	}

	cfg, ok := configs[conn]
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid connection: %s", conn))
	}
	driver := fmt.Sprintf(
		"host=%s port=%d dbname=%s search_path=%s user=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s connect_timeout=%d",
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SearchPath,
		cfg.Username,
		cfg.Password,
		cfg.SslMode,
		cfg.SslCert,
		cfg.SslKey,
		cfg.SslRootCert,
		cfg.ConnectTimeout,
	)
	db, err := gorm.Open("postgres", driver)
	if err != nil {
		return nil, err
	}

	if config.AppModeIs(gin.DebugMode) || config.AppModeIs(gin.TestMode) {
		db.LogMode(true)
		//db.SetLogger(log.New(os.Stdout, "\n", 0))
	} else {
		db.LogMode(false)
	}

	db.DB().SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)

	db.SingularTable(true)
	dbs[conn] = db
	return db, err
}
