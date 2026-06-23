//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a program demonstrating how slice capacity grows automatically when appending.
*/

func main() {
	s := []int{}

	for i := range 11 {
		s = append(s, i)
		fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	}
}
