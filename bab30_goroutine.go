package main

import (
	"fmt"
	"runtime"
)

func printing(data int, message string) {
	for i := 0; i < data; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	// * Using concurently go with using runtime
	go printing(5, "halo")
	// * Normal
	printing(5, "apa kabar")

	// * Input data from user
	var input string
	fmt.Printf("Please enter the value: ")
	fmt.Scanln(&input)
	fmt.Println("Testing using Scanln and the Output was =", input)
}
