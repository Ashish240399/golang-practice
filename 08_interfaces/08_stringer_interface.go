//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Implement the fmt.Stringer interface on a custom struct.
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name=%s, Age=%d", p.Name, p.Age)
}
func main() {
	p := Person{
		Name: "Ashish",
		Age:  34,
	}

	fmt.Println(p)
}
