//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Write a variadic function that calculates the sum of an arbitrary number of integers.
*/

func main() {
	ans := sum(2, 4, 5, 9)

	fmt.Println(ans)
}

func sum(num ...int) int {
	total := 0

	for _, n := range num {
		total += n
	}

	return total
}
