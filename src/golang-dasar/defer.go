package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil Function")
}

func runApplication(value int) {
	fmt.Println("Running application")
	result := 10 / value
	fmt.Println("Result: ", result)
	logging()

}

func main() {
	runApplication(0)
}
