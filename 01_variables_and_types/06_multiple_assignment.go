//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Use multiple assignment to swap the values of two variables.
*/

func main() {
	// Write your solution here
	a := 10
	b := 30

	a, b = b, a
	fmt.Println(a, b)
}
