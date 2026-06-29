//go:build ignore

package main

import (
	"errors"
	"fmt"
)

/*
Difficulty: Easy

Question: Write a function that returns an error using errors.New.
*/

func getError() error {
	return errors.New("This is an error")
}

func main() {
	err := getError()

	if err != nil {
		fmt.Println(err)
	}
}
