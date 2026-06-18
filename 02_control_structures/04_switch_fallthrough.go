//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Demonstrate the use of the fallthrough keyword in a switch statement.
*/

func main() {
	const age = 9

	switch{
	case age > 18:
		fmt.Println("You are eligible")
		fallthrough
	case age <= 18:
		fmt.Println("You are not eligible")
		fallthrough
	case age < 0:
		fmt.Println("You are just a seed")
	}

}
