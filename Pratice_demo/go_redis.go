package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping().Result()
	return err
}
func main() {
	if err := initClient(); err != nil {
		fmt.Printf(" init  redis failed =====> 🚀🚀🚀 %v\n", err)
	}
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {

		}
	}(rdb)
}
