//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Demonstrate variable shadowing in an inner scope.
*/

func main() {
	name := "Ashish"

	if true {
		name := "XYZ"
		fmt.Println(name)
	}

	fmt.Println(name)
}
