package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.Now()

	switch {
	case time.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("Its's after noon")

	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool type")
		case int:
			fmt.Println("int type")
		default:
			fmt.Println("I don't know what it is")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("you don't know")
}
