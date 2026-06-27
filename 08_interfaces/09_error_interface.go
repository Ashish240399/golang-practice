//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Implement the error interface on a custom struct.
*/

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("{Field:%s, Message:%s}", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 18 {
		return ValidationError{
			Field:   "Age",
			Message: "You are not allowed",
		}
	}
	return nil
}

func main() {
	age := 13

	err := validateAge(age)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You are eligible")
	}
}
