package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Initialize the Redis client

	// e71dae798c95  conatiner id of redis running in same network of this golang conatiner. 
	rdb := redis.NewClient(&redis.Options{
		Addr:     "e71dae798c95:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Default DB
	})

	// Check connection to Redis
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("could not connect to redis: %v", err)
	}
	fmt.Println("Connected to Redis")

	// Set a key in Redis
	err = rdb.Set(ctx, "mykey", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatalf("could not set key: %v", err)
	}
	fmt.Println("Set key 'mykey' with value 'Hello, Redis!'")

	// Get the value of a key from Redis
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatalf("could not get key: %v", err)
	}
	fmt.Printf("The value of 'mykey' is: %s\n", val)
}
