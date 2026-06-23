//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Create a slice from an array and modify an element through the slice.
*/

func main() {
	arr := [5]int{23, 45, 67, 76, 43}

	slice := arr[:4]

	slice[0] = 11

	fmt.Println(slice)
}
