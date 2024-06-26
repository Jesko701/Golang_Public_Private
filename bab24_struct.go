package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   string
	grade int
}

// * Anonymous Struct
var s3 = struct {
	person
	hobby string
}{
	person: person{"John", 20},
	hobby:  "Reading",
}

func main() {
	// * Struct Definition
	var p1 = person{name: "wick", age: 21}
	var s1 = student{person: p1, grade: 2}

	// * If using pointer, then the change of variable s2 will be affect to the original (s1)
	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)

	fmt.Println("student 3, name :", s3.name)
	fmt.Println("student 3, name :", s3.hobby)

}
