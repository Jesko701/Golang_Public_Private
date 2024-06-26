package main

import "fmt"

func main() {
	// Ini adalah komentar
	fmt.Println("Hello World")
	/* testing comment
	123 */

	var firstName string = "Jovial"
	var lastName string = "Pattiasina"
	angkatan := 2020

	fmt.Printf("halo %s %s! angkatan %d \n", firstName, lastName, angkatan)

	// deklarasi multivariable
	/*
		var satu, dua, tiga string = "jason", "jacob", "taichi"
		one, two, three := "lord", 123, "in the heavers"

	*/
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f", decimalNumber)
	// deklarasi menggunakan keyword
	name := new(string)
	fmt.Println(*name)

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
	fmt.Printf("angka %d dan %d menggunakan angka, sedangkan untuk %s dan %s menggunakan string", satu, dua, three, four)
}
