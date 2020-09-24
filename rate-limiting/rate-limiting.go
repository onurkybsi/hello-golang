package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.

func main() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// By blocking on a receive from the limiter channel before serving each request,
		// we limit ourselves to 1 request every 200 milliseconds.
		<-limiter

		fmt.Println("request", req, time.Now())
	}
}
