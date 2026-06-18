//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Declare a custom type based on an int and write a function using it.
*/

type Age int

func main() {
	
	printAge := func(age Age){
		fmt.Println(age)
	}

	printAge(23)
}
