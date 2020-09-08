package main

import "fmt"

// intSeq() returns the integer returning anonymous function.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	// This function value captures its own i value, which will be updated each time we call nextInt.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()

	fmt.Println(nextInt2())
	fmt.Println(nextInt2())

	// Notice ! This print 4
	fmt.Println(nextInt())
}
