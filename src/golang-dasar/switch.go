package main

import (
	"fmt"
)

func main() {
	name := "Joko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi,Siapa Kamu?")
	}

	fmt.Println("---------switch short expression-------------")

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// fmt.Println(string(65))
}
