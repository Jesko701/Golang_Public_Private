package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	println()

	var i = 1
	for {
		fmt.Printf("%d, ", i)
		i++
		if i == 5 {
			break
		}
	}
	println()
	// loop()
	// loopBreakContinue()
	// nestedLoop()
	loopLabel()
}

func loop() {
	println()
	fmt.Println("Versi Looping Golang")
	var xs = "123"
	for i, v := range xs {
		fmt.Println("Index", i, "Value=", v)
	}

	fmt.Println("Menggunakan panjang array: ")
	var ys = [5]int{10, 20, 30, 40, 50} //array
	for _, v := range ys {
		fmt.Print(v, " ")
	}

	fmt.Println("Menggunakan slice: ")
	var zs = ys[0:2] //slice
	for _, v := range zs {
		fmt.Print(v, " ")
	}

	fmt.Println("Menggunakan map: ")
	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	for k, v := range kvs {
		fmt.Println("Key", k, "Value", v)
	}
	for range kvs {
		fmt.Println("Menggunakan range tanpa variable")
	}

	// menentukan nilai numerik perulangan
	fmt.Println("Menggunakan banyaknya nilai numerik perulangan: ")
	for i := range 5 {
		fmt.Print(i, " ")
	}
}

// penggunaan break dan continue
func loopBreakContinue() {
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}

// perulangan bersarang
func nestedLoop() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

// label perulangan
func loopLabel() {
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks, [", i, "][", j, "]", "\n")
		}
	}
}
