package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var Сtx = context.Background()
var Client *redis.Client

func RedisInit() {
	LoadEnv()
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_addr") + ":" + os.Getenv("redis_port"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})

	pong, err := Client.Ping(Сtx).Result()
	if err != nil {
		log.Fatalln("Error connection to redis:", err)
	}
	fmt.Println("Connected to redis", pong)
}
