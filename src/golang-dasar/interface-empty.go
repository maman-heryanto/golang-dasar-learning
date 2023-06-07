package main

import "fmt"

//interface kosong return apapun
//type data menjadi interface{}
func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {

	// var data int = ups(1) //salah var data
	var data interface{} = ups(1)
	fmt.Println(data)
}
