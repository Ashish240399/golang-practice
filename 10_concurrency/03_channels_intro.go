//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a channel, send a value into it from a goroutine, and receive it in main.
*/

func sendMessage(message string, ch chan string) {
	ch <- message
}

func main() {
	ch := make(chan string)

	go sendMessage("Hello", ch)

	message := <-ch

	fmt.Println(message)
}
