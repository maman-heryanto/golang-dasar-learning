package main

import "fmt"

func main() {
	name1 := "Eko"
	name2 := "Budi"

	result := name1 == name2
	fmt.Println(result)

	value1 := 100
	value2 := 200

	fmt.Println("---------------")
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
