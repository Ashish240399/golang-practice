//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Declare a pointer to an integer, print its address and value.
*/

func main() {
	a:=10

	b := &a
	fmt.Println(b, *b)
}
