package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `json:"mode" ini:"mode"`
	Port         int    `json:"port" ini:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
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
	Host         string `json:"host" ini:"host"`
	Password     string `json:"password" ini:"password"`
	Port         int    `json:"port" ini:"port"`
	DB           int    `json:"db" ini:"db"`
	PoolSize     int    `json:"pool_size" ini:"pool_size"`
	MinIdleConns int    `json:"min_idle_conns" ini:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `json:"level" ini:"level"`
	Filename   string `json:"filename" ini:"filename"`
	MaxSize    int    `json:"max_size" ini:"max_size"`
	MaxAge     int    `json:"max_age" ini:"max_age"`
	MaxBackups int    `json:"max_backups" ini:"max_backups"`
}

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./conf/")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
