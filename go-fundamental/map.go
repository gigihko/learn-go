package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person ["name"] = "Gigih"
	// person ["address"] = "BSN"

	person := map[string]string{
		"name":    "Gigih",
		"address": "BSN",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Gigih"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
