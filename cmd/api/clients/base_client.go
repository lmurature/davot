package clients

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

const (
	redisHost = "cache"
	redisPort = "6379"
)

var (
	CacheClient cacheClient
)

func InitializeClients() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	res, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from redis: ", res)
	CacheClient = cacheClient{redisClient}
}
