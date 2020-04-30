package settings

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

type AppConfig struct {
	*ServerConfig `json:"server" ini:"server"`
	*LogConfig    `json:"log" ini:"log"`
	*MySQLConfig  `json:"mysql" ini:"mysql"`
	*RedisConfig  `json:"redis" ini:"redis"`
}

type ServerConfig struct {
	Port int `json:"port" ini:"port"`
}

type MySQLConfig struct {
	Host         string `json:"host" ini:"host"`
	User         string `json:"user" ini:"user"`
	Password     string `json:"password" ini:"password"`
	DB           string `json:"db" ini:"db"`
	Port         int    `json:"port" ini:"port"`
	MaxOpenConns int    `json:"max_open_conns" ini:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns" ini:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `json:"host" ini:"host"`
	Password string `json:"password" ini:"password"`
	Port     int    `json:"port" ini:"port"`
	DB       int    `json:"db" ini:"db"`
}

type LogConfig struct {
	Level      string `json:"level" ini:"level"`
	Filename   string `json:"filename" ini:"filename"`
	MaxSize    int    `json:"max_size" ini:"max_size"`
	MaxAge     int    `json:"max_age" ini:"max_age"`
	MaxBackups int    `json:"max_backups" ini:"max_backups"`
}

func LoadFromFile(cfgFile string) (err error) {
	return ini.MapTo(Conf, cfgFile)
}

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return err
}
