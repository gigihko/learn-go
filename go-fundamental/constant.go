package main

import "fmt"

func main() {
	const (
		firstName string = "Gigih"
		LastName         = "Pambuko"
	)

	fmt.Println(firstName)
	fmt.Println(LastName)
	//error
	// firstName = "Budi"
	// LastName = "Joko"
}
