package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// use select with a default clause to implement non-blocking sends.
	// message and signals have not any value so default case will immediately take.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
