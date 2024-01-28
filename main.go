package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	pool := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     10,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Minute,
	})
	fmt.Println(pool)
	defer pool.Close()
	setErr := pool.Set(context.Background(), "name", "tom cat", 0)
	if setErr != nil {
		fmt.Println(setErr)
	}
	val, getErr := pool.Get(context.Background(), "name").Result()
	if getErr != nil {
		fmt.Println(getErr)
	}
	fmt.Println(val)
}
