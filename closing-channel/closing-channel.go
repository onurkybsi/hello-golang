package main

import (
	"fmt"
)

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received new job", j)
			} else {
				fmt.Println("there is no job that received")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}

	// thats enough sends, close the channel
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
