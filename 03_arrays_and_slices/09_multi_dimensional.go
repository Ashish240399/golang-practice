//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Create and initialize a 2D slice (slice of slices) and print its contents.
*/

func main() {
	var matrix [][]int

	matrix = append(matrix, []int{1, 2})

	fmt.Println(matrix)
}
