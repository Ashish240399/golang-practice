//go:build ignore

package main

import (
	"fmt"
	"sort"
)

/*
Difficulty: Hard

Question: Write a function that prints the keys of a map in alphabetical sorted order.
*/

func main() {
	data := map[string]int{
		"banana": 3,
		"apple":  1,
		"orange": 2,
		"grape":  4,
	}

	keys := make([]string, 0)

	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, el := range keys {
		fmt.Println(el)
	}
}
