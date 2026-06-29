//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Create a custom error struct that implements the error interface.
*/

type ErrorStruct struct {
	Message string
}

func (e ErrorStruct) Error() string {
	return e.Message
}

func main() {
	errorData := ErrorStruct{
		Message: "This is error message",
	}

	fmt.Println(errorData.Error())
}
