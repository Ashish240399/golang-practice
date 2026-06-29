//go:build ignore

package main

import (
	"fmt"
	"time"
)

/*
Difficulty: Easy

Question: Start a basic goroutine that prints a message concurrently.
*/

func execute() {
	fmt.Println("Hello execution")
}

func main() {
	go execute()

	fmt.Println("Hello 2")

	time.Sleep(time.Second)
}
