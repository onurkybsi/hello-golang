package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	id  int
	res chan int
}

type writeOp struct {
	id  int
	val int
	res chan bool
}

func main() {
	var readOpCount uint64
	var writeOpCount uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.res <- state[read.id]

			case write := <-writes:
				state[write.id] = write.val
				write.res <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			read := readOp{
				id:  rand.Intn(5),
				res: make(chan int)}
			reads <- read

			<-read.res

			atomic.AddUint64(&readOpCount, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			write := writeOp{
				id:  rand.Intn(5),
				val: rand.Intn(100),
				res: make(chan bool)}
			writes <- write

			<-write.res

			atomic.AddUint64(&writeOpCount, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOpCount)
	fmt.Println(readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOpCount)
	fmt.Println(writeOpsFinal)
}
