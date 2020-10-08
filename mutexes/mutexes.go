package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readCounter uint32
	var writeCounter uint32

	for r := 0; r < 100; r++ {
		go func() {
			readed := 0
			for {
				someKey := rand.Intn(5)
				mutex.Lock()
				readed += state[someKey]
				mutex.Unlock()

				atomic.AddUint32(&readCounter, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				someKey := rand.Intn(5)
				writingVal := rand.Intn(100)
				mutex.Lock()
				state[someKey] = writingVal
				mutex.Unlock()

				atomic.AddUint32(&writeCounter, 1)
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println(atomic.LoadUint32(&readCounter))
	fmt.Println(atomic.LoadUint32(&writeCounter))
}
