package main

import (
	"fmt"
	"strings"
)

func main() {
	// * this is a function Varadic that's hold lot of data simultaniously
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3} // * slice (bagian dari array dan lebih fleksible)
	var avg = calculate(numbers...)
	fmt.Printf("Rata-rata : %.2f\n", avg)
	var hobby = []string{"ngoding", "golang", "C#"}
	// * user input a lot of data
	paramAndVariadic("Joy", "makan", "tidur", "maen")
	// * insert data slice to varadict function
	paramAndVariadic("Joy_2", hobby...)
}

func calculate(angka ...int) float64 {
	var total int = 0
	for _, v := range angka {
		total += v
	}
	rata_rata := float64(total) / float64(len(angka))
	return rata_rata
}

func paramAndVariadic(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
