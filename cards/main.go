package main

import "fmt"

type contact struct {
	phone   string
	zipCode int
}

type person struct {
	firstName      string
	lastName       string
	contactDetails contact
}

func main() {
	listFruits := fruits()

	for _, fruit := range listFruits {
		fmt.Println(fruit)
	}

	newPerson := person{firstName: "Ashish", lastName: "Mohanty", contactDetails: contact{
		phone:   "2342342",
		zipCode: 32423,
	}}
	newPersonPointer := &newPerson

	newPerson.updateName("Rupali")
	changeLastName("Jaiswal", &newPerson)

	fmt.Printf("%+v", newPerson)
	fmt.Printf("%+v", newPersonPointer)
}

func fruits() []string {

	fruits := []string{"apple", "orange", "Lichi"}

	return fruits
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func changeLastName(newLastName string, pointerToPerson *person) {
	pointerToPerson.lastName = newLastName
}
