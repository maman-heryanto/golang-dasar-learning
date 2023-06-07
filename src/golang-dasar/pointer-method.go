package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	// fmt.Println("Hello, Mr.", man.Name)
	man.Name = "Mr " + man.Name
}

func main() {
	eko := Man{"eko"}
	eko.married()

	fmt.Println(eko.Name)
}
