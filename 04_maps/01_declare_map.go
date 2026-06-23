//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Declare and initialize a map with string keys and integer values.
*/

func main() {
	profile := map[string]int{}

	profile["Ashish"] = 23

	fmt.Println(profile)
}
