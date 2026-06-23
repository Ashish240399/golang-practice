//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Iterate over a map using for-range and print all keys and values.
*/

func main() {
	profile := map[string]int{
		"Ashish": 23,
		"Rupali": 21,
	}

	for key, value := range profile {
		fmt.Printf("Key=%s, Value=%d\n", key, value)
	}
}
