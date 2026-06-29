//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

/*
Difficulty: Easy

Question: Start an HTTP server listening on port 8080 returning 'Hello World' for '/'.
*/

func normalRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", normalRoute)

	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error:", err)
		return
	}

}
