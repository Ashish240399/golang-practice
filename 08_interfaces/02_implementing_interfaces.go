//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Have multiple different structs implement the same interface.
*/

type Person interface {
	GetName() string
	GetAge() int
}

type Men struct {
	name string
	age  int
}

type Tree struct {
	name string
	age  int
}

func (m Men) GetName() string {
	return m.name
}

func (t Tree) GetAge() int {
	return t.age
}

func main() {
	men := Men{
		name: "Ashish",
		age:  45,
	}

	tree := Tree{
		name: "Bannian",
		age:  99,
	}

	fmt.Println("Men name: ", men.GetName())
	fmt.Println("Tree age: ", tree.GetAge())
}
