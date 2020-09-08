package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1

	fmt.Println("Initial value : ", i)

	zeroval(i)
	fmt.Println("Processed with zeroval : ", i)

	zeroptr(&i)
	fmt.Println("Processed with zeroptr : ", i)

	fmt.Println("Address of i: ", &i)
}
