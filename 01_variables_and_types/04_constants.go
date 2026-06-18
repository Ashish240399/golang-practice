//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

/*
Difficulty: Easy

Question: Declare untyped and typed constants and use them in a calculation.
*/

func main() {
	// Write your solution here
	const untypedVariable = 10

	const typedVariable string = "The age of the child is: "

	res := typedVariable + strconv.Itoa(untypedVariable)
	fmt.Println(res)
}
