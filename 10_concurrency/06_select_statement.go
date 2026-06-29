//go:build ignore

package main

import (
	"fmt"
	"time"
)

/*
Difficulty: Hard

Question: Use a select statement to wait on multiple channel operations.
*/

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello chan 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello chan 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
