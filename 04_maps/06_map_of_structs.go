//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a map where the values are custom structs.
*/

func main() {
	type customStruct struct {
		name string
		age  int
	}

	data := map[string]customStruct{
		"Candidate 1": {
			name: "Ashish",
			age:  23,
		},
	}

	fmt.Println(data)
}
