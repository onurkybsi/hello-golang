package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// mechanism for managing state in Go is communication over channels :
// sync/atomic package for atomic counters accessed by multiple goroutines.

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				// To atomically increment the counter we use AddUint64
				atomic.AddUint64(&ops, 2)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops: ", ops)
}
