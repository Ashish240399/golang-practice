//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Write a function that accepts an empty interface (any type).
*/

func printVal(v any) {
	fmt.Println(v)
}

func main() {
	printVal(23)
	printVal("Ashish")
	printVal(true)
}
