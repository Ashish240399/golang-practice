//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Use defer to ensure a file is closed (simulate file close with print).
*/

func main() {
	fmt.Println("File opened")
	defer fmt.Println("File closed")
	fmt.Println("Reading file")
}
