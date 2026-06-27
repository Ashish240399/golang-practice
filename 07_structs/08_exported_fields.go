//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Demonstrate the difference between exported (capitalized) and unexported fields.
*/
type Person struct {
	Name string // Public
	age  int    // Private
}

func main() {
	data := Person{
		Name: "Ashish",
		age:  56,
	}

	fmt.Println(data.GetAge())

}

func (p Person) GetAge() int {
	return p.age
}
