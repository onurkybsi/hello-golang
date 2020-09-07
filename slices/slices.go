package main

import "fmt"

func main() {

	// To create an empty slice with non-zero length, use the builtin make()
	s := make([]string, 3)
	fmt.Println("Empty slice : ", s)

	s[0] = "element1"
	s[1] = "element2"
	s[2] = "element3"

	// Length of slice
	fmt.Println(len(s))

	// append(), which returns a slice containing one or more new values
	s = append(s, "newElement1", "newElement2")
	fmt.Println(s)

	s2 := make([]string, len(s))

	// copy slice s to slice s2
	copy(s2, s)
	fmt.Println("Copy slice of s, s2: ", s2)

	// slice operator ":" like Python
	s3 := s[2:5]
	fmt.Println("This get s[2], s[3], s[4]: ", s3)

	s3 = s[2:]
	fmt.Println("This get from s[2] to last element: ", s3)

	// We can initialize slices like this:
	s4 := []string{"str1", "str2", "str3"}
	fmt.Println(s4)

	// 2d slices:
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		// 2. dimension length
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
