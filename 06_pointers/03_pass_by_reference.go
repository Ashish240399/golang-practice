//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Write a function that modifies a variable passed by pointer.
*/

func main() {
	a := 12

	pointerVal := &a

	modifyByPointer(pointerVal)

	fmt.Println(a)
}

func modifyByPointer(data *int) {
	*data = 20
}
