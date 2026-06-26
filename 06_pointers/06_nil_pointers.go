//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Check if a pointer is nil before dereferencing it to avoid panics.
*/

func main() {
	a := 10

	b := &a

	if b != nil {
		fmt.Println(*b)
	} else {
		fmt.Println("Pointer is nil")
	}
}
