//go:build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
Difficulty: Hard

Question: Write a simple middleware function that logs the request method and path.
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Path: %s\n", r.URL.Path)
		fmt.Printf("Body: %s\n", string(body))

		// Restore the body
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
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
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	loggedMux := loggerMiddleware(mux)

	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", loggedMux)

	if err != nil {
		fmt.Println("Server error: ", err)
	}
}
