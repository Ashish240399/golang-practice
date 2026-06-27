//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Compare two structs for equality and explain when structs are comparable.
*/

// In Go, structs can be compared using == and !=, but only if all of their fields are comparable.

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Name string
	Age  int
	// Height int
}

func main() {
	person := Person{
		Name: "Ashish",
		Age:  23,
	}

	student := Student{
		Name: "Ashish",
		Age:  23,
	}

	if person == Person(student) {
		fmt.Println("All are equal")
	} else {
		fmt.Println("Not equal")
	}
}
