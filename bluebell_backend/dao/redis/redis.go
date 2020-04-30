package redis

import (
	"bluebell_backend/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Client.Close()
}
