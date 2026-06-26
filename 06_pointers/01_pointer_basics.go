//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Declare a pointer, assign it a variable's address, and print the dereferenced value.
*/

func main() {
	a := 10
	b := &a

	fmt.Println(b)

	fmt.Println(*b)

}
