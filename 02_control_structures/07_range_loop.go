//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Use a for-range loop to iterate over a string and print index and character.
*/

func main() {
	words:= "Golang tutorial"

	for _, ch:=range words{
		fmt.Printf("%c\n",ch)
	}
}
