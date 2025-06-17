package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Ayra"
	names[1] = "Aghniya"
	names[2] = "Pambuko"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[2] = "Budi"

	fmt.Println(names[2])

	var values = [...]int{
		90,
		81, 
		80,
		100,
		110
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))
	// fmt.Println(len(values))

}
