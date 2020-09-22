package main

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources or that
// otherwise need to bound execution time.

func main() {

	channel := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "result"
	}()

	// timeout pattern :
	// When res doesnt receive by res in 1 second, second case available.
	select {
	case res := <-channel:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
