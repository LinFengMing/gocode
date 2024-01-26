package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()
	fmt.Println(rdb)
	err := rdb.HSet(context.Background(), "user01", "name", "john").Err()
	if err != nil {
		fmt.Println("Error setting hash name:", err)
		return
	}
	err = rdb.HSet(context.Background(), "user01", "age", 10).Err()
	if err != nil {
		fmt.Println("Error setting hash age:", err)
		return
	}
	result, err := rdb.HGet(context.Background(), "user01", "name").Result()
	if err == redis.Nil {
		fmt.Println("Field not found in the hash.")
	} else if err != nil {
		fmt.Println("Error getting hash name:", err)
	} else {
		fmt.Printf("Value of %s in %s: %s\n", "name", "user01", result)
	}
	result2, err := rdb.HGet(context.Background(), "user01", "age").Result()
	if err == redis.Nil {
		fmt.Println("Field not found in the hash.")
	} else if err != nil {
		fmt.Println("Error getting hash age:", err)
	} else {
		fmt.Printf("Value of %s in %s: %s\n", "age", "user01", result2)
	}
	fieldsValues := map[string]interface{}{
		"name": "tom",
		"age":  18,
	}
	err = rdb.HMSet(context.Background(), "user02", fieldsValues).Err()
	if err != nil {
		fmt.Println("Error setting multiple hash fields:", err)
		return
	}
	fieldsToGet := []string{"name", "age"}
	results, err := rdb.HMGet(context.Background(), "user02", fieldsToGet...).Result()
	if err != nil {
		fmt.Println("Error getting multiple hash fields:", err)
		return
	}

	for i, field := range fieldsToGet {
		fmt.Printf("Value of %s in %s: %v\n", field, "user02", results[i])
	}
}
