//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

/*
Difficulty: Medium

Question: Read query parameters from an incoming GET request.
*/

func handler(w http.ResponseWriter, r *http.Request) {
	paramName := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Name is = %s", paramName)
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error: ", err)
	}
}
