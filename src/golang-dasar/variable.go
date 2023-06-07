package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)
	name = "Eko Khenddy"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age int8 = 30
	fmt.Println(age)

	coutry := "Indonesia"
	fmt.Println(coutry)

	coutry = "America"
	fmt.Println(coutry)

	fmt.Println("----------------")

	var (
		firstName = "Eko"
		lastName  = "Kurniawan"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
