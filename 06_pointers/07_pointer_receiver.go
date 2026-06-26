//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a method with a pointer receiver to modify the struct's state.
*/

type Person struct {
	name string
	age  int
}

func main() {
	personData := Person{
		name: "Ashish",
		age:  45,
	}

	personData.structModifier()

	fmt.Println(personData)
}

func (pointer *Person) structModifier() {
	pointer.age = 33
}
