//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

/*
Difficulty: Easy

Question: Use http.HandleFunc to route different paths to different handler functions.
*/

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handler 1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handler 2")
}

func main() {

	http.HandleFunc("/1", handler1)

	http.HandleFunc("/2", handler2)

	fmt.Println("Server is running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error: ", err)
	}

}
