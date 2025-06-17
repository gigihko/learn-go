package main

import "fmt"

func main() {

	type NoKTP string

	var ktpGiko NoKTP = "1111111111"

	var contoh string = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpGiko)
	fmt.Println(contohKtp)
}
