package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a parent context with a 5 second deadline.
	parentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a child context with a value.
	childCtx := context.WithValue(parentCtx, "key", "value")

	// Simulate a long-running process via a goroutine.
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println(time.Now(), ctx.Value("key"))
			}
		}

	}(childCtx)

	// Block the main function to allow go routine to finish.
	time.Sleep(10 * time.Second)
}
