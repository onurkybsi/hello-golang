package main

import "fmt"

// By default channels are unbuffered,
// meaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values
// without a corresponding receiver for those values.

// Unbuffered channels block receivers until data is available on the channel and senders
// until a receiver is available. Buffered channels will only block a sender once the buffer fills up.

func main() {
	// a channel of string buffering up to 2 values
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
