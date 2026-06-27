//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a struct that contains another struct as a field.
*/

type Address struct {
	pinCode string
}

type Person struct {
	name    string
	age     int
	address Address
}

func main() {
	userData := Person{
		name: "Ashish",
		age:  56,
		address: Address{
			pinCode: "87689",
		},
	}

	fmt.Println(userData)
}
