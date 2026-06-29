//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write functions that accept send-only and receive-only channels.
*/

func sendOnlyChan(ch chan<- string) {
	ch <- "Hello from sendOnly"
}

func receiveOnlyChan(ch <-chan string) {
	message := <-ch

	fmt.Println(message)
}

func main() {
	ch := make(chan string)
	go sendOnlyChan(ch)
	receiveOnlyChan(ch)
}
