//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a function that returns a closure which keeps track of a counter state.
*/

func main() {
	inc := counter()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
