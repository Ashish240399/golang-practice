//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Write a function that takes two integers and returns their sum.
*/

func main() {
	ans := returnSum(3, 5)

	fmt.Println(ans)
}

func returnSum(a int, b int) int {
	return a + b
}
