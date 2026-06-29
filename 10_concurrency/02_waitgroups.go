//go:build ignore

package main

import (
	"fmt"
	"sync"
)

/*
Difficulty: Medium

Question: Use sync.WaitGroup to wait for multiple goroutines to finish.
*/

func printMessage(message string, wg *sync.WaitGroup) {
	fmt.Println(message)

	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	printMessage("First goroutine", &wg)
	printMessage("Second goroutine", &wg)
	printMessage("Third goroutine", &wg)

	wg.Wait()

	fmt.Println("All finished")
}
