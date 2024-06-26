package main

import "fmt"

type person struct {
	name string
	age  int
}

var allStudents = []person{
	{"wick", 22},
	{"ethan", 21},
	{"bourne", 25},
}

func main() {
	for _, students := range allStudents {
		fmt.Println(students.name, "age", students.age)
	}
}
