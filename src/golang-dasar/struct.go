package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct func
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello,", name, "My name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Subang"
	eko.Age = 20
	fmt.Println(eko)

	eko.sayHi("Joko")

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(joko)
}
