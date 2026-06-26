//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a function that returns a pointer to a local variable (escape analysis).
*/

func main() {
	ans := a()

	fmt.Println(ans)
}

func a() *int {
	val := 20

	return &val
}
