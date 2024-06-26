package main

import "fmt"

// * Anonymous Struct with Slice

type person struct {
	name string
	age  int
}

func main() {
	allStudents := []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, students := range allStudents {
		fmt.Println(students)
	}
}
