package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println("--------------------------------")

	// months[5] = "BulanNew"
	// fmt.Println(slice1)

	// slice1[0] = "MeyAgain"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Bulan_New")
	fmt.Println(slice3)
	slice3[1] = "now_december"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	fmt.Println("--------------------------------")
	// REPLACE Slice
	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	fmt.Println(days)
	daySlice := days[5:]
	daySlice[0] = "Saturday_New"
	daySlice[1] = "Sunday_New"
	fmt.Println(days)

	fmt.Println("-----------Make Slice----------")
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("-------------Copy Slice------------")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	fmt.Println("--------------------------------")

	iniArrays := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArrays)
	fmt.Println(iniSlice)
}
