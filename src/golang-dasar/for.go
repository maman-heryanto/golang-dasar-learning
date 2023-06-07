package main

import (
	"fmt"
)

func main() {

	for counter := 1; counter < 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
		// log.SetPrefix("LOG: ")
		// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
		// log.Println("init started")
	}

	slice := []string{"Eko", "Kurniawan", "Khannedy"}

	fmt.Println("----------for-------------------")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println("----------for range-------------")
	for i, names := range slice {
		fmt.Println("index", i, "=", names)
	}

	fmt.Println("----------for range di map------")
	person := make(map[string]string)
	person["name"] = "Eko"
	person["age"] = "22"
	person["address"] = "Subang"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
