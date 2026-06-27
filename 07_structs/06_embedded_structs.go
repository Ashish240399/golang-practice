//go:build ignore

package main

import "fmt"

/*
Difficulty: Hard

Question: Demonstrate struct embedding (composition) to simulate inheritance.
*/

type Animal struct {
	name string
}

func (a Animal) sleep() {
	fmt.Println(a, " is sleeping")
}

type Dog struct {
	Animal
	LegCount int
}

func (d Dog) bark() {
	fmt.Println("Only dog can bark")
}

func main() {
	parentAnimal := Animal{
		name: "Cow",
	}

	parentAnimal.sleep()

	childAnimal := Dog{
		Animal: Animal{
			name: "Dog",
		},
		LegCount: 4,
	}

	childAnimal.bark()

	childAnimal.sleep()
}
