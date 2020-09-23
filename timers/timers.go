package main

import (
	"fmt"
	"time"
)

// Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// Channell of timer = C
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// One reason a timer may be useful is that you can cancel the timer before it fires.

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stopped := timer2.Stop()

	if stopped {
		fmt.Println("Timer 2 stopped. Thats good thing")
	}

	time.Sleep(2 * time.Second)
}
