//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Demonstrate the LIFO (Last-In, First-Out) execution order of multiple defers.
*/

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}
