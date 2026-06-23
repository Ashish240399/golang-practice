//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Use the make built-in function to create a slice with a specific length and capacity.
*/

func main() {
	a := make([]int, 3, 5)

	a = append(a, 3)
	fmt.Println(a)
}
