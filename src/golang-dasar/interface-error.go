package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak Boleh Noll")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := pembagian(100, 0)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Hasil", hasil)
}
