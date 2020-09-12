package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// running it synchronously.
	f("direct")

	// This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	// start goroutine for an anonymous func
	go func(msg string) {
		fmt.Println(msg)
	}("going with new goroutine")

	time.Sleep(time.Second)
	fmt.Println("done!")
}
