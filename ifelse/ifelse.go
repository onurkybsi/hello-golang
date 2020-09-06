package main

import "fmt"

func main() {
	if 11%2 == 0 {
		fmt.Println("11 is even")
	} else {
		fmt.Println("11 is odd")
	}

	if varDeclared := 1; varDeclared > 0 {
		fmt.Println(varDeclared, " is positive")
	} else if varDeclared == 0 {
		fmt.Println(varDeclared, " is zero")
	} else {
		fmt.Println(varDeclared, " is negative")
	}
}
