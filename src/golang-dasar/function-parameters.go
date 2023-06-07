package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello to", firstName, lastName)
}

func main() {
	sayHelloTo("Eko", "Kurniawan")
}
