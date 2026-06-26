//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Change the value of a variable using its pointer.
*/

func main() {
	a := 10

	b := &a

	*b = 30

	fmt.Println(a)
}
