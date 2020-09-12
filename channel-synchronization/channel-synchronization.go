package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.

// The done channel will be used to notify another goroutine
// that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)

	go worker(done)

	// done ı receive(get yaptıgımız icin) asagıyı bloklamis oluyoruz.
	// main thread worker func ın bitmesini bekliyor
	<-done

	// If you removed the <- done line from this program,
	// the program would exit before the worker even started.

	fmt.Println("This execute after done == true")
}
