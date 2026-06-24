//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Write a function that returns both the quotient and remainder of division.
*/

func main() {
	q, r := division(10, 3)

	fmt.Println(q, r)
}

func division(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}
