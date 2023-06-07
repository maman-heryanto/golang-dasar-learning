package main

import "fmt"

func main() {

	fmt.Println("-------------if------------------")
	name := "Eko"
	if name == "Eko" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Hi,Boleh Kenalan " + name)
	}

	fmt.Println("-------------else if--------------")
	name = "aseppp"

	if name == "Eko" {
		fmt.Println("Hello " + name)
	} else if name == "joko" {
		fmt.Println("Hi,Boleh Kenalan " + name)
	} else {
		fmt.Println("Bukan Eko/Joko,Boleh Kenalan " + name)
	}

	fmt.Println("-------------if Short Statement--------------")
	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Name Benar")
	}
}
