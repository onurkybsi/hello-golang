package main

import (
	"fmt"
	"time"
)

// Default unbuffered

func main() {
	channel := make(chan string, 1)

	// buffer = 1 so sender goroutine in this case following func doesnt block
	// if we add channel <- "test2" to following func, buffer will be filled so func will be blocked
	go func() {
		channel <- "test"

		fmt.Println("Send!")
	}()

	go func() {
		time.Sleep(3 * time.Second)

		fmt.Println(<-channel)
	}()

	time.Sleep(4 * time.Second)
	fmt.Println("Exited!")
}
