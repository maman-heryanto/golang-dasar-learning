package main

import "fmt"

func getGoodBye(name string) string {
	return "Good by " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("eko"))
}
