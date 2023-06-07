package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("eko kurniawan", "eko"))
	fmt.Println(strings.Contains("eko kurniawan", "budi"))

	fmt.Println(strings.Split("Eko Kurniawan Khenndy", " "))

	fmt.Println(strings.ToLower("Eko Kurniawan Khenndy"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Khenndy"))
	fmt.Println(strings.ToTitle("Eko Kurniawan Khenndy"))

	fmt.Println(strings.Trim("    Eko Kurniawan   ", " "))
	fmt.Println(strings.ReplaceAll("Eko Eko Eko Eko Eko Eko Eko Eko Eko ", "Eko", "Budi"))

}
