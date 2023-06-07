package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "The host to connect to database")
	var user *string = flag.String("user", "root", "The user to connect to database")
	var password *string = flag.String("password", "toor", "The password to connect to database")

	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("User", *user)
	fmt.Println("Password", *password)
}
