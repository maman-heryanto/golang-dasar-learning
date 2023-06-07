package main

import "fmt"

func getFullName() (string, string, string) {
	return "Eko", "Kurniawan", "khandedly"
}
func main() {
	firstName, middleName, lastName := getFullName() // jika variable tidak di gunakan underscore(_)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
