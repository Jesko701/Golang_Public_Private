package main

import "fmt"

type person struct {
	name string
	age  int
}
type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Printf(
		"name		: %s\n"+
			"age		: %d\n"+
			"person age	: %d\n"+
			"grade		: %d\n",
		s1.name, s1.age, s1.person.age, s1.grade,
	)
	// * Anonymous Struct
	var s2 = struct {
		person
		grade int
	}{}
	s2.person = person{"fredy", 21}
	s2.grade = 2
	fmt.Printf(
		"\nname		: %s\n"+
			"age		: %d\n"+
			"person age	: %d\n"+
			"grade		: %d\n",
		s2.name, s2.age, s2.person.age, s2.grade,
	)
}
