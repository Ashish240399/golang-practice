//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Define and invoke an anonymous function inside main.
*/

func main() {
	randomFunc := func() {
		fmt.Println("Hello random function")
	}

	randomFunc()

	func() {
		fmt.Println("I am IIFE")
	}()
}
