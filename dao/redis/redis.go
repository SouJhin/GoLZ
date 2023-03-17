package redis

import (
	"fmt"

	"server/settings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	return err
}

func Close() {
	_ = rdb.Close()
}
