//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Use a type switch to handle different concrete types stored in an interface.
*/

func main() {
	var data any = 23

	switch v := data.(type) {
	case string:
		fmt.Println("Data is string", v)

	case int:
		fmt.Println("Data is int", v)

	default:
		fmt.Println("Data is unknown", v)

	}
}
