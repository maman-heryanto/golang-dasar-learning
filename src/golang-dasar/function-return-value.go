package main

import "fmt"

func getHello(name string) string {
	// return "Hello " + name
	if name == "" {
		return "Hello Bro"
	}
	return "Hello " + name
}

func main() {
	resutl1 := getHello("Eko kurniawan")
	resutl2 := getHello("")
	fmt.Println(resutl1)
	fmt.Println(resutl2)
}
