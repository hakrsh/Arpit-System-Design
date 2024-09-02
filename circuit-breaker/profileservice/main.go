package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	for {
		// Simulating service going down and coming back up
		status := "up"
		if rand.Intn(100) < 30 {
			status = "down"
		}
		err := rdb.Publish(ctx, "profile_service_status", status).Err()
		if err != nil {
			log.Printf("Error publishing to Redis: %v\n", err)
		}
		fmt.Printf("Published status: %s\n", status)
		time.Sleep(5 * time.Second)
	}
}
