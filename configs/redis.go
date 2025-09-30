package configs

import (
	 "github.com/go-redis/redis/v8"
	 "fmt"
)

func NewRedisClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:   "localhost:6379",
        Password: "",
        DB:       0,
    })
fmt.Println("success connect to redis")
    return client
}