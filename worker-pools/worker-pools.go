package main

import (
	"fmt"
	"time"
)

// we’ll run several concurrent instances with worker-pool pattern

func worker(id int, jobs <-chan int, results chan<- int) {
	// receive works with jobs channel
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)

		// representation of some work
		time.Sleep(time.Second)

		// send result of job with results job
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5

	// send jobs to worker pool with jobs channel
	jobs := make(chan int, numJobs)
	// collect results from worker pool results channel
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
