package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(host)

	//set
	os.Setenv("APP_USERNAME", "root@localhost")
	os.Setenv("APP_PASSWORD", "toor")
	//get
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}
