package persistence

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

var Redis *redis.Client

func MustGetRedisClient() *redis.Client {
	if Redis != nil {
		return Redis
	}

	options, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		fmt.Printf("error parsing redis url (%s): %v\n", os.Getenv("REDIS_URL"), err)
		panic(err)
	}

	client := redis.NewClient(options)
	Redis = client
	return client
}
