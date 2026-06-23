//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Add new key-value pairs to a map and update existing ones.
*/

func main() {
	profile := map[string]int{
		"Ashish": 23,
		"Rupali": 21,
	}

	profile["Ashish"] = 24

	fmt.Println(profile)
}
