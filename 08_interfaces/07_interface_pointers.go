//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Demonstrate the difference between a value receiver and a pointer receiver when implementing interfaces.
*/

type Person interface {
	Walk()
}

type Men struct {
	name string
}

func (m Men) Walk() {
	fmt.Println("Hello value rcv func:", m.name)
}

func main() {
	var s Person

	data := Men{
		name: "Ashish",
	}
	s = data
	s.Walk()

	ptrData := &Men{
		name: "John",
	}
	s = ptrData
	s.Walk()

}
