package main

import "fmt"

func main() {
	// * interface not restricted by type of data
	var secret interface{}
	secret = "Jovial Pattiasina"
	fmt.Println(secret)
	secret = []string{"Eat", "Code", "Cloud Platform"}
	fmt.Println(secret)
	secret = 12.4
	fmt.Println(secret)

	// * map
	var data map[string]interface{}
	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)

	// * Null interface
	type person struct {
		name string
		age  int
	}
	var orang interface{} = &person{name: "wick", age: 24}
	var name = orang.(*person).name
	fmt.Println(name)

	// * Key Value using map[] and interface{}
	var keyValue = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}
	fmt.Println()
	for _, each := range keyValue {
		fmt.Println(each["name"], "age is", each["age"])
	}

}
