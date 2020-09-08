package main

import "fmt"

func vals() (string, string) {
	return "returned value 1", "returned value 2"
}

func main() {
	val1, val2 := vals()

	fmt.Println(val1)
	fmt.Println(val2)

	// ignore returned first value
	_, val22 := vals()
	fmt.Println(val22)
}
