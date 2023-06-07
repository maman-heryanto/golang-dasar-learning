package main

import "fmt"

func main() {
	type NoKTP string
	type Merried bool

	var noKTPMaman NoKTP = "1231232132132"
	var merriedStatus Merried = true
	fmt.Println(noKTPMaman)
	fmt.Println(merriedStatus)
}
