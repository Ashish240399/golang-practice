//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Write a higher-order function that takes a function as an argument and executes it.
*/

func main() {
	ans := hof(func(a int, b int) int {
		return a + b
	})

	greet := func() {
		fmt.Println("Hello Ann")
	}

	execute(greet)

	fmt.Println(ans)
}

func hof(fn func(a int, b int) int) int {
	return fn(10, 20)

}

func execute(fn func()) {
	fn()
}
