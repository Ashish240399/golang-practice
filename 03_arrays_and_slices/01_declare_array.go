//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Declare an array of 5 integers, initialize it, and print the values.
*/

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, e := range arr {
		fmt.Println(e)
	}
}
