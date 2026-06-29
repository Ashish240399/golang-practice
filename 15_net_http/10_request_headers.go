//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

/*
Difficulty: Medium

Question: Read specific headers from a request and set custom headers in the response.
*/

func headerHandler(w http.ResponseWriter, r *http.Request) {
	// Read request headers
	userAgent := r.Header.Get("User-Agent")
	authToken := r.Header.Get("Authorization")

	// Set custom response headers
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-App-Name", "Go HTTP Server")
	w.Header().Set("X-Version", "1.0")

	fmt.Fprintf(w, "User-Agent: %s\n", userAgent)
	fmt.Fprintf(w, "Authorization: %s\n", authToken)
}

func main() {
	http.HandleFunc("/", headerHandler)

	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
