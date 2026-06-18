//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

/*
Difficulty: Easy

Question: Convert an integer to a float64 and a string to an integer.
*/

func main() {
	// Write your solution here
	var a int = 10
	var b string = "20"

	var ax float64 = float64(a)
	bx, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println(ax, bx)
}
