//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a sub-slice from an existing slice (slicing a slice).
*/

func main() {
	s := []int{3, 4, 5, 6, 7, 8, 9, 0, 12, 2, 33}

	newS := s[3:7]
	fmt.Println(newS)
}
