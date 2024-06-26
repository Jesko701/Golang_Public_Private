package main

import "fmt"

func main() {
	alternativeArray()
	println()
	dynamicArray()
	println()
	multiDimensionArray()
}

func conventionalArray() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "g"
	names[2] = "water"
	names[3] = "law"
	fmt.Println(names[0], names[1], names[2], names[3])
}

func alternativeArray() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Print("panjang dari fruits ", len(fruits), " : ")
	for _, v := range fruits {
		fmt.Print(v, " ")
	}
}

// jumlah dinamis array tanpa mengetahui panjang array
func dynamicArray() {
	var numbers = [...]int{2, 3, 4, 5}
	fmt.Print("Ini adalah perulangan print array tanpa mengetahui panjang array: ")
	for _, v := range numbers {
		fmt.Print("Angka ke ", v, " - ")
	}

}

func multiDimensionArray() {
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
}
