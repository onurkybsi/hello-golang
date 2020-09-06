package main

import "fmt"

const constVal string = "constant"

func main() {
	fmt.Println(constVal)

	const constVal2 = 3.14
	fmt.Println(constVal2)
}
