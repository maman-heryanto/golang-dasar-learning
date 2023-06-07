package main

import "fmt"

func random() interface{} {
	// return "string"
	// return 100
	return false
}

func main() {
	//type assertions
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)

	}
}
