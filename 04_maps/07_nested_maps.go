//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Create and interact with a nested map (map of maps).
*/

func main() {
	parentMap := make(map[string]map[string]int)

	parentMap["Ashish"] = map[string]int{
		"Age": 29,
	}

	parentMap["Ashish"]["Age"] = 34

	fmt.Println(parentMap)
}
