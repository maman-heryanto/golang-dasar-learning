package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHolla(hasName HasName) {
	fmt.Println("Holla", hasName.GetName())
}

//PERSON
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

//ANIMAL
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	sayHolla(eko)

	cat := Animal{
		Name: "Cat",
	}
	sayHolla(cat)

}
