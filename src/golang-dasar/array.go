package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khenddy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("-----------")
	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}
