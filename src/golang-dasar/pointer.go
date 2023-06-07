package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	fmt.Println("------------1------------")
	addressTest1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	addressTest2 := addressTest1 //type data *address

	addressTest2.City = "Bandung"

	fmt.Println(addressTest1)
	fmt.Println(addressTest2)

	fmt.Println("-------------2--------------")
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //type data *address
	var address3 *Address = &address1

	address2.City = "Bandung"

	address2 = &Address{"Malang", "Jawa timur", "Indonesia"} //address1 tidak berubah
	// *address2 = Address{"Malang", "Jawa timur", "Indonesia"} //address1 mengikuti address2 karena terdapat * pada address2

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println("-----------")

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	address4.Provincy = "DKI"
	address4.Country = "Indonesia"
	fmt.Println(address4)

	alamat := Address{
		City:     "Jakarta",
		Provincy: "DKI",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}
