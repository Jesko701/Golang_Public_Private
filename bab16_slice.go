package main

import "fmt"

func main() {
	/*
		var fruitsA = []string{"apple", "grape"}     // slice
		var fruitsB = [2]string{"banana", "melon"}   // array
		var fruitsC = [...]string{"papaya", "grape"} // array

		fmt.Println(fruitsA[0:2])
		fmt.Println(fruitsB[0:1]) // Hanya akan print index 0 karena rumusnya [index awal : panjang data - 1]
		fmt.Println(fruitsC[0:2])
	*/
	// nambahSlice()
	copySlice()
}

func sliceData() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
}

func nambahSlice() {
	var fruits = []string{"apple", "grape", "banana"}
	var bfruits = fruits[0:2]
	fmt.Println(cap(bfruits)) // 3
	fmt.Println(len(bfruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bfruits) // ["apple", "grape"]

	var cFruits = append(bfruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bfruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]
}

func copySlice() {
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	m := copy(dst, src)
	fmt.Println(m)   // 3
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange

	dst_2 := []string{"potato", "potato", "potato"}
	src_2 := []string{"watermelon", "pinnaple"}
	n := copy(dst_2, src_2)
	fmt.Println(dst_2) // watermelon pinnaple potato
	fmt.Println(src_2) // watermelon pinnaple
	fmt.Println(n)     // 2
}

func slice3indeks() {
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
