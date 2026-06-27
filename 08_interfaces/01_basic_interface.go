//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Define an interface with a single method and a struct that implements it.
*/

type Person interface {
	GetAge() int
}

type User struct {
	Name string
	Age  int
}

func (u User) GetAge() int {
	return u.Age
}

func main() {
	var p Person
	p = User{
		Name: "Ashish",
		Age:  34,
	}

	fmt.Println("User's age: ", p.GetAge())
}
