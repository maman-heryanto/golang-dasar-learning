package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "John",
		"age":     "22",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])    //map[key]
	fmt.Println(person["age"])     //map[key]
	fmt.Println(person["address"]) //map[key]

	fmt.Println("-------Function MAP---------")
	fmt.Println(len(person)) //len(map)

	fmt.Println("----------------------------")
	books := make(map[string]string)
	books["title"] = "Golang Programming"
	books["author"] = "Eko Kurniawan"
	books["genre"] = "Learning"

	fmt.Println(books)
	fmt.Println(len(books))

	delete(books, "genre")
	fmt.Println(books)
	fmt.Println(len(books))

	fmt.Println("----------------------------")

}
