package main

import (
	"fmt"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	circuitOpen := false
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Subscribe to profile service status
	sub := rdb.Subscribe(ctx, "profile_service_status")
	defer sub.Close()

	// Start a goroutine to listen for status updates
	go func() {
		ch := sub.Channel()
		for msg := range ch {
			if msg.Payload == "down" {
				circuitOpen = true
				fmt.Println("Profile service is down, circuit breaker tripped.")
			} else {
				circuitOpen = false
				fmt.Println("Profile service is up, circuit breaker closed.")
			}
		}
	}()

	// Simulating requests from post service
	for {
		if !circuitOpen {
			fmt.Println("Making call to profile service...")
			// Simulate the profile service call here
			time.Sleep(2 * time.Second) // Simulating delay between calls
		} else {
			fmt.Println("Circuit breaker is open. Skipping call to profile service.")
			time.Sleep(3 * time.Second) // Retry after a short delay
		}
	}
}
