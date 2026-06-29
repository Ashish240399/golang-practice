//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Difficulty: Medium

Question: Write an HTTP handler that returns a JSON-encoded struct with Content-Type set.
*/

type Greeting struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	greetingData := Greeting{
		Message: "Hello",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(greetingData)
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error: ", err)
	}
}
