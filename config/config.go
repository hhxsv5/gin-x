package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	AppCfg   map[string]interface{}
	BasePath string
	cfgName  string
)

func init() {
	flag.StringVar(&BasePath, "path", ".", "the base path of project")
	flag.StringVar(&cfgName, "config", "config", "the name of toml configuration file")
	flag.Parse()

	viper.AddConfigPath(BasePath)
	viper.AddConfigPath(".")
	viper.SetConfigName(cfgName)
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	AppCfg = viper.GetStringMap("app")
	log.Println(fmt.Sprintf("Load config successfully! Config path: %s/%s.toml", BasePath, cfgName))
}

func AppMode() string {
	t, ok := AppCfg["mode"].(string)
	if !ok {
		return gin.ReleaseMode
	}
	return t
}

func AppModeIs(mode string) bool {
	return AppMode() == mode
}

func AppCron() bool {
	t, ok := AppCfg["cron"].(bool)
	if !ok {
		return false
	}
	return t
}

func AppListenAddress() string {
	l, ok := AppCfg["listen"].(string)
	if !ok {
		return "127.0.0.1:5200"
	}
	return l
}
