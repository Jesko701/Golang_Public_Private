package main

import "fmt"

func main() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Println(isEqual)
}
