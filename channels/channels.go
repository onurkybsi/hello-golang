package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels
// from one goroutine and receive those values into another goroutine.

func main() {

	// Create a new channel with make(chan val-type)
	message := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the message channel we made above,
	// from a new goroutine.
	go func() {
		time.Sleep(time.Second)
		message <- "ping"
	}()

	// The <-channel syntax receives a value from the channel.
	// Here we’ll receive the "ping" message we sent above
	msg := <-message

	fmt.Println(msg)
	// below statement thrown fatal error: all goroutines are asleep - deadlock!
	// above goroutine are asleep after msg := <-message statement
	// degeri aktardıktan sonra thread ölüyor gibi. Tekrar alamıyoruz
	fmt.Println(<-message)
	fmt.Println("awaited for ping above")

	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program
	// for the "ping" message without having to use any other synchronization.
}
