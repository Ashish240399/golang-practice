//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Initialize a map with a specific capacity hint using make.
*/

func main() {
	m := make(map[string]int, 100)

	m["Ashish"] = 23

	fmt.Println(m)
}
