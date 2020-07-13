package main

import "fmt"

func main() {
	//verbs
	// %v -> general default format
	var verb = "verb"
	fmt.Printf("This is the %v like in C", verb)

	fmt.Println()

	var verb2 = "verb2"
	fmt.Printf("This is the type verb, so type of verb2 is: %T", verb2)

	fmt.Println()

	var verb3 = 1
	fmt.Printf("This is the int verb: %d", verb3)

	fmt.Println()

	var verb4 = 0.2
	fmt.Printf("This is the float verb: %.2f", verb4)

	fmt.Println()

	//Sprint and Sprintln, Sprintf

	var string1 = "This is string.Format() "
	var string2 = "in Go."

	var resultString = fmt.Sprint(string1, string2)
	fmt.Print(resultString)

	var resultString2 = fmt.Sprintf("%v(advanced) %v", string1, string2)
	fmt.Println(resultString2)

	//Getting User Input
	var valueFromUser string

	fmt.Scan(&valueFromUser)

	fmt.Printf("This is the value from user: %v", valueFromUser)
}
