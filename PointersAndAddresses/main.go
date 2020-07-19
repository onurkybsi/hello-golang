package main

import "fmt"

func testFunc(num int) {
	num += 100
}

func testFunc2(num2 *int) {
	*num2 += 100
}

func main() {

	//
	num := 1
	testFunc(num)
	fmt.Println(num)

	//
	addressOfString := "Test"
	fmt.Println(&addressOfString)

	//
	testString := "TestString"
	var goPointer *string = &testString
	fmt.Println(goPointer)

	//dereferencing or indirecting
	firstValue := "Value1"
	point := &firstValue
	// change the value of address:
	*point = "Value2"
	fmt.Println(firstValue)

	//
	num2 := 2
	testFunc2(&num2)
	fmt.Println(num2)
}
