package main

import "fmt"

func sum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func subtract(num1 int, num2 int) int {
	return num1 - num2
}

func main() {
	res1 := sum(3, 4, 5)
	fmt.Println("Result 1 : ", res1)

	res2 := subtract(10, 2)
	fmt.Println("Result 2 : ", res2)
}
