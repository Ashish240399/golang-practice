//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Use the copy built-in function to copy elements from one slice to another.
*/

func main() {
	s := []int{43, 4456, 56, 545, 454, 5}

	c := make([]int, len(s))

	n := copy(c, s)

	fmt.Println(n)
	fmt.Println(c)
}
