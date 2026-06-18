//go:build ignore

package main

import "fmt"

/*
Difficulty: Easy

Question: Write a basic if-else statement to check if a number is even or odd.
*/

func main() {
	num:=11

	if (num%2 == 1){
		fmt.Println("This is ODD")
	}else{
		fmt.Println("This is EVEN")
	}
}
