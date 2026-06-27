//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create and use an anonymous struct in a single line.
*/

func main() {
	randomStruct := struct {
		name string
		age  int
	}{
		name: "Ashish",
		age:  33,
	}

	fmt.Print(randomStruct)
}
