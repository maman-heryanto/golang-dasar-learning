package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// var person map[string]string = nil
	// fmt.Println(person)

	person1 := NewMap("Eko")
	fmt.Println(person1)

	person2 := NewMap("")
	fmt.Println(person2)
	if person2 == nil {
		fmt.Println("No person")
	}

}
