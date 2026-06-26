//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Declare a pointer to a pointer and trace the values back to the original variable.
*/

func main() {
	val := 12

	ptrVal := &val

	secondPtrVal := &ptrVal

	fmt.Println(val, ptrVal, secondPtrVal, *ptrVal, **secondPtrVal)
}
