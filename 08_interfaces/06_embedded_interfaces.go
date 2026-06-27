//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Create an interface that embeds other interfaces.
*/

type Human interface {
	Walk() string
}

type Men interface {
	Run() string
}

type Person interface {
	Human
	Men
}

type Employee struct{}

func (e Employee) Walk() string {
	return ("I am walking")
}

func (e Employee) Run() string {
	return ("I am running")
}

func main() {
	var person Person = Employee{}

	fmt.Println(person.Run())
	fmt.Println(person.Walk())
}
