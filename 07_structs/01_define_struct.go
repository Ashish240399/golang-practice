//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Define a struct representing a User with Name, Age, and Email.
*/
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	user := User{
		Name:  "Ashish",
		Age:   23,
		Email: "ewihf@gmail.com",
	}

	fmt.Println(user)
}
