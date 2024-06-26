package main

import "fmt"

type Person struct {
	name string
	age  int
}

type People = Person

func main() {
	var p1 = struct {
		name string
		age  int
	}{age: 22, name: "wick"}
	var p2 = struct {
		name string
		age  int
	}{"ethan", 23}

	fmt.Println(p1, p2)

	fmt.Println("casting Alias type dari objek")
	people := People{"wick", 21}
	fmt.Println(Person(people))

	person := Person{"wick", 21}
	fmt.Println(People(person))

}
