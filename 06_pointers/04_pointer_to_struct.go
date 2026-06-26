//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a pointer to a struct and access its fields.
*/
type Person struct {
	name string
	age  int
}

func main() {
	a := &Person{
		name: "",
		age:  0,
	}

	a.name = "efgdjvh"
	a.age = 38

	fmt.Println(*a)
}
