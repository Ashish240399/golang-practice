//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a buffered channel and send multiple values without a receiver block.
*/

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- 40
	fmt.Println(<-ch)

}
