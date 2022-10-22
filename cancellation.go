package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	run := func(ctx context.Context) {
		n := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exiting")
				close(ch)
				return
			default:
				time.Sleep(time.Millisecond * 300)
				fmt.Println(n)
				n++
			}

		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("goodbye")
		cancel()
	}()

	go run(ctx)

	fmt.Println("waiting to cancel....")

	<-ch

	fmt.Println("bye")
}
