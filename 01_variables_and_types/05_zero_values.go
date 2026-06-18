//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Print the zero values of various built-in types.
*/

func main() {
	// Write your solution here
	var a int
	var b float32
	var c float64
	var d string
	var e struct {}
	var f bool

	fmt.Println(a, b, c, d, e, f)
}
