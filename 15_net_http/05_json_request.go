//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Difficulty: Hard

Question: Write a handler that decodes a JSON request body into a struct.
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User data: %+v", user)
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error: ", err)
	}
}
