//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Extract the underlying concrete value from an interface using type assertion.
*/

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{
		name: "Ashish",
		age:  34,
	}

	var value any = person

	x, ok := value.(Person)

	if ok == true {
		fmt.Println(x.name)
	} else {
		fmt.Println("Data is no of type person")
	}
}
