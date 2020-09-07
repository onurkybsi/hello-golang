package main

import "fmt"

func main() {
	var arr1 [5]int

	fmt.Println("Empty int arr1ay with default int values : ", arr1)

	arr1[4] = 1
	fmt.Println("Array1 v2 : ", arr1)
	fmt.Println("Array1 v2 5th element: ", arr1[4])

	// length of array
	fmt.Println(len(arr1)) // = 5

	// Another array declaration
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
