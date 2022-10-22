package main

import (
	"context"
	"fmt"
	"time"
)

const shortduration = 1 * time.Second

// Context is used to provide timeout or cancellation for go routine
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), shortduration)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("slept")

	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}
}
