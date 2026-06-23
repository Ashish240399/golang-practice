//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Iterate over an array using a for loop and print the sum of its elements.
*/

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	sum := 0

	for _, el := range arr {
		sum += el
	}

	fmt.Println((sum))
}
