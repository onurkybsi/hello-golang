package main

import "fmt"

func main() {
	// range iterates over elements in a variety of data structures.

	nums := []int{2, 3, 4}
	sum := 0

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	// range with map
	kv := map[string]string{"k1": "v1", "k2": "v2"}

	for key, value := range kv {
		fmt.Println("Key : "+key+" Value : ", value)
	}

	// range on strings iterates over Unicode code points.
	for index, character := range "go" {
		fmt.Println(index, character)
	}
}
