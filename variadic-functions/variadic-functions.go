package main

import "fmt"

// Variadic functions can be called with any number of arguments.
func sum(nums ...int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

func main() {
	sumOfDigits := sum(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Also, Println is a common variadic function.
	fmt.Println("Sum of Digits : ", sumOfDigits)

	numsAsArray := []int{1, 2, 3, 4, 5}
	fmt.Println("Array parameter for variadic functions : ", sum(numsAsArray...))
}
