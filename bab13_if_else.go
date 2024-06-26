package main

import "fmt"

func main() {
	var point = 8840.0

	// * using if, also temp variable
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// * switch case for lot condition (seems same like java but actually not)
	pointer := 6
	switch pointer {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Note Bad")
			fmt.Println("You can do it better")
		}
	}

	pointer_2 := 9
	switch {
	case pointer_2 > 10:
		fmt.Println("Perfect")
	case (pointer_2 >= 5) && (pointer_2 <= 10):
		fmt.Println("Good")
	default:
		{
			fmt.Println("Silahkan dicoba lagi")
			fmt.Println("Kamu pasti bisa")
		}
	}
}
