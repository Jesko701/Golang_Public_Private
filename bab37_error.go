package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// * recover
func catch() {
	if i := recover(); i != nil {
		fmt.Println("Error occured:", i)
	} else {
		fmt.Println("Apps run perfectly")
	}
}

func main() {
	// * Postponed the execution until it runs all
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}
}
