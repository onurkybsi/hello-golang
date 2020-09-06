package main

import "fmt"

func main() {
	var a = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Print default value. In this case = 0
	var e int
	fmt.Println(e)

	// declare and initializing a variable
	// The following equal to var f string = "test"
	f := "test"
	fmt.Println(f)
}
