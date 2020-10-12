package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{9, 5, 12, 20, 6}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println(sort.IntsAreSorted(ints))
}
