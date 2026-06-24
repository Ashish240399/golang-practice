//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a recursive function to calculate the factorial of a number.
*/

func main() {
	fact := factorial(5)

	fmt.Println(fact)
}

func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n-1)
}
