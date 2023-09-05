package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pingResult, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("PING:", pingResult)

	_, err = redisClient.Incr(ctx, "count").Result()
	if err != nil {
		panic(err)
	}

	count, err := redisClient.Get(ctx, "count").Int()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Count: %d\n", count)

	err = redisClient.Close()
	if err != nil {
		panic(err)
	}
}
