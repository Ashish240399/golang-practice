//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Create a slice containing pointers to integers and iterate over it.
*/

func main() {
	a := 12
	b := 34

	sliceData := []*int{
		&a, &b,
	}

	for _, data := range sliceData {
		fmt.Println(data, *data)
	}
}
