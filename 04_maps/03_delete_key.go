//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Delete an element from a map using the delete function.
*/

func main() {
	profile := map[string]int{
		"Ashish": 23,
		"Rupali": 21,
	}

	delete(profile, "Ashish")

	fmt.Println(profile)
}
