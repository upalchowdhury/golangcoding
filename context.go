package main

import (
	"context"
	"fmt"
	"time"
)

const shortduration = 1 * time.Second

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
