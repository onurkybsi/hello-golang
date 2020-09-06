package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j%2 == 0)
	}

	for {
		fmt.Println("Just one iteration")
		break
	}

	for k := 0; k <= 5; k++ {
		if k%2 == 0 {
			continue
		}

		fmt.Println(k)
	}
}
