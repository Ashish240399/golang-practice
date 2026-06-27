//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

/*
Difficulty: Medium

Question: Add JSON struct tags to a struct and briefly explain their purpose.
*/

type Person struct {
	FirstName string `json:"first_name"`
}

func main() {
	data := Person{
		FirstName: "Ashish",
	}

	newData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(newData))
}
