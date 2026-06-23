//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Append multiple elements to a slice and print the new slice and its capacity.
*/

func main() {
	a := make([]int, 0, 2)

	a = append(a, 2, 3, 4, 5, 6)

	fmt.Println(a, len(a), cap(a))
}
