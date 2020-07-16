package main

import "fmt"

func sayHello() {

	fmt.Println("Hello !")
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func goCanReturnTwoValues(isDivisibleByTwo int) (bool, int) {
	boolResult := false
	intResult := 0

	if isDivisibleByTwo%2 == 0 {
		boolResult = true
		intResult = 1
	}

	return boolResult, intResult
}

func queryDatabase(query string) string {
	var result string
	connectDatabase()
	//Go to run a function, but at the end of the current function(scope): DEFER
	defer disconnectDatabase()

	if query == "SELECT * FROM Person;" {
		result = "NAME|SURNAME\nOnur|Kayabasi"
	}

	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}

func main() {
	sayHello()

	result := sum(4, 5)

	fmt.Println(result)

	fmt.Println(goCanReturnTwoValues(8))
	fmt.Println(goCanReturnTwoValues(7))

	queryDatabase("SELECT * FROM Person;")
}
