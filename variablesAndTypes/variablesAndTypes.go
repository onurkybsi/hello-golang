package main

import "fmt"

func main() {

	//named const values:
	const namedConstValue = "This is the named constant value"
	fmt.Println(namedConstValue)

	//named variable values:
	var integerData int = 4
	fmt.Println(integerData * 2)

	var isGoPretty bool
	isGoPretty = true
	fmt.Println(isGoPretty)

	var stringVariable string = "Go programming language is the best"
	fmt.Println(stringVariable)

	//default values of variables:
	var defaultValueOfInt int
	var defaultValueOfString string
	var defaultValueOfbool bool

	fmt.Println(defaultValueOfInt, defaultValueOfString, defaultValueOfbool)

	//Variable definition without specifying type

	boolVariableWithoutSpecifyingType := true
	intVariableWithoutSpecifyingType := 5
	stringVariableWithoutSpecifyingType := "Variable definition without specifying type"

	fmt.Println(boolVariableWithoutSpecifyingType, intVariableWithoutSpecifyingType, stringVariableWithoutSpecifyingType)

	// OR

	var anotherBoolVariableWithoutSpecifyingType = true
	var anotherIntVariableWithoutSpecifyingType = 5
	var anotherStringVariableWithoutSpecifyingType = "Variable definition without specifying type"

	fmt.Println(anotherBoolVariableWithoutSpecifyingType, anotherIntVariableWithoutSpecifyingType, anotherStringVariableWithoutSpecifyingType)

	//Multiple Variable Declaration

	var string1, string2 string
	string1 = "String1"
	string2 = "String2"
	fmt.Println(string1, string2)

	integer, boolean := 4, true
 	fmt.Println(integer, boolean)
}
