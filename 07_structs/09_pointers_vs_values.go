//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Show the difference between copying a struct by value vs passing by pointer.
*/

type Person struct {
	Name string
	Age  int
}

func updateByValue(p Person) {
	p.Age = 23

	fmt.Println("Inside updateByValue: ", p)
}

func updateByPointer(p *Person) {
	p.Age = 23

	fmt.Println("Inside updateByPointer: ", p)
}

func main() {
	person := Person{
		Name: "Ashish",
		Age:  67,
	}

	fmt.Println("Person before update by value: ", person)
	updateByValue(person)
	fmt.Println("Person after update by value: ", person)

	fmt.Println("Person before update by pointer: ", person)
	updateByPointer(&person)
	fmt.Println("Person after update by pointer: ", person)
}
