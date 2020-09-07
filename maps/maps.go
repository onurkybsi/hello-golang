package main

import "fmt"

func main() {
	// maps are like dictionaries in c#

	// creating map
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map m : ", m)

	v1 := m["k1"]
	fmt.Println(v1)

	// get length of map
	fmt.Println(len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k1")
	fmt.Println(m)
}
