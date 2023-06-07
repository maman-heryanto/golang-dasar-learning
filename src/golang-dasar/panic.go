package main

import "fmt"

func endApp() {
	fmt.Println("Application Stopped...")
	message := recover() // recover from panic
	fmt.Println("Terjadi error: ", message)
}

func runApp(error bool) {
	defer endApp() //defer
	if error {
		panic("ERROR") //panic
	}
	fmt.Println("Application Run")
}

func main() {
	runApp(true)
	fmt.Println("Eko") //masih di eksekusi meskipun error di func run app
}
