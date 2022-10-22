package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {

	time.Sleep(time.Second)
	return rand.Intn(100)

}

func main() {
	//
	dataChan := make(chan int) //can be a custom type like type newtype string
	// can add buffer for temp storage

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Println(n)
	}

	//n := <-dataChan // it does not have temp storage need to enter and exit at the same time

}
