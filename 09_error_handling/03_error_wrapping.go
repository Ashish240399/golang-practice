//go:build ignore

package main

import (
	"errors"
	"fmt"
)

/*
Difficulty: Medium

Question: Use fmt.Errorf with the %%w verb to wrap an error.
*/

func getError() error {
	return errors.New("This is error")
}

func main() {
	err := getError()

	if err != nil {
		wrappedErr := fmt.Errorf("Error: %w", err)
		fmt.Println(wrappedErr)
	}
}
