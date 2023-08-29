package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Incr(ctx, "count").Result()
	if err != nil {
		panic(err)
	}

	count, err := rdb.Get(ctx, "count").Int()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Count: %d\n", count)
}
