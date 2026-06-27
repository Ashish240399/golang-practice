//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a factory function (e.g., NewUser) that returns a pointer to a new struct.
*/

type Person struct {
	Name  string
	Age   int
	Email string
}

func NewUser(name string, age int, email string) *Person {
	user := Person{
		Name:  name,
		Age:   age,
		Email: email,
	}

	return &user
}

func main() {
	newUser := NewUser("Ashish", 45, "sdh@gmail.com")

	data := *newUser

	fmt.Println(data)
	fmt.Println(data.Age, newUser.Age)
}
