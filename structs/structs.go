package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}

	p.age = 23

	return &p
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	alice := person{name: "Alice", age: 30}

	fmt.Println(alice)
	fmt.Println(alice.name)
	fmt.Println(alice.age)

	aliceptr := &alice

	fmt.Println(aliceptr.name)
	fmt.Println(aliceptr.age)

	// new struct with newPerson func
	fmt.Println(newPerson("Jon"))
}
