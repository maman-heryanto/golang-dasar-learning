package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		name := "Budi" // not closure
		fmt.Println(name)

		fmt.Println("increment")
		counter++
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
