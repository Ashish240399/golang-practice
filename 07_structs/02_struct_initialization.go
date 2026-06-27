//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Initialize a struct using field names and print it.
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
