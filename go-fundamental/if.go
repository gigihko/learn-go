package main

import "fmt"

func main() {
	name := "kurniawan"

	if name == "Gigih" {
		fmt.Println("Hello Gigih")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hai, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
