//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Check if a specific key exists in a map using the comma-ok idiom.
*/

func main() {
	profile := map[string]int{
		"Ashish": 23,
		"Rupali": 21,
	}

	age, ok := profile["Ashish"]

	if ok {
		fmt.Println("Ashish", age)
	} else {
		fmt.Println("Key does not exist")
	}
}
