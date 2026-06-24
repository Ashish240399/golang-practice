//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Rewrite the previous function using named return values.
*/

func main() {
	q, r := division(10, 3)

	fmt.Println(q, r)
}

func division(a int, b int) (quotient, remainder int) {

	quotient = a / b
	remainder = a % b

	return
}
