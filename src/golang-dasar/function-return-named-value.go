package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Barker"
	lastName = "Hall"
	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
