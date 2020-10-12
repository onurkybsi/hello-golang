package main

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go, we need a corresponding type.
// Here we’ve created a byLength type that is just an ALIAS for the builtin []string type.
type sortedByLength []string

// We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function.
func (s sortedByLength) Len() int {
	return len(s)
}

func (s sortedByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(sortedByLength(fruits))
	fmt.Println(fruits)
}
