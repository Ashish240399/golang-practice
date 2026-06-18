//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Write a switch statement to print the day of the week based on an integer.
*/

func main() {

	day:=1

	switch day {
	case 1,7:
		fmt.Println("Weekend")
	case 2,3,4,5,6:
		fmt.Println("Weekdays")
	}
}
