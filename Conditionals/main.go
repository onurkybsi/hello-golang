package main

import "fmt"

func main() {
	var boolean = false

	if boolean {
		fmt.Println("Thats true")
	} else {
		fmt.Println("Thats false")
	}

	rightTime := true
	rightPlace := true

	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	readyToGo := true

	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}

	num := 14

	if num%6 == 0 {
		fmt.Println("It can be divided by 6")
	} else if num%3 == 0 {
		fmt.Println("It can be divided by 3")
	} else if num%3 == 0 {
		fmt.Println("It can be divided by 2")
	} else {
		fmt.Println("It can not be divided by 6, 3 or 2")
	}

	switch num % 2 {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	//Scoped Short Declaration Statement
	x := 8
	y := 9
	if product := x * y; product > 60 {
		fmt.Println(product, "  is greater than 60")
	}

	switch product := x * y; product > 60 {
	case true:
		fmt.Println(product, "  is greater than 60")
	default:
		fmt.Println(product, "  is smaller than 60")
	}
}
