package librarytutorial

import "fmt"

func init() {
	p1 := Student{}
	p1.Name = "Jesko"
	p1.grade = 22
	fmt.Println("--> librarytutorial/library.go imported")
}

func SayHello() {
	fmt.Println("Ini adalah fungsi public")
}

func introduce(name string) {
	fmt.Println("nama saya adalah", name)
}
