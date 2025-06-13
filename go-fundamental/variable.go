package main

import "fmt"

func main() {
	// var name string //ini opsional, karena golang bisa tau tipe datanya

	name := "Ayra Pambuko" //contoh deklarasi variable pengganti "var" adalah ":="
	fmt.Println(name)

	name = "Ayra Aghniya"
	fmt.Println(name)

	var (
		firstName  = "Ayra"
		middleName = "Aghniya"
		LastName   = "Pambuko"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(LastName)
}
