package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	Charset         string
	Collation       string
	Loc             string
	ParseTime       string `mapstructure:"parse_time"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
}

func init() {
	err := viper.UnmarshalKey("databases.mysql", &configs)
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
			cfg.Port = 3306
		}
		if cfg.Charset == "" {
			cfg.Charset = "utf8mb4"
		}
		if cfg.Collation == "" {
			cfg.Collation = "utf8mb4_general_ci"
		}
		if cfg.Loc == "" {
			cfg.Loc = "Local"
		}
		if cfg.ParseTime == "" {
			cfg.ParseTime = "false"
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
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&loc=%s&parseTime=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		cfg.Collation,
		cfg.Loc,
		cfg.ParseTime,
	)
	db, err := gorm.Open("mysql", driver)
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
