package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var areaMain, kelilingMain = luasMultipleReturn(diameter)
	fmt.Printf("Luas lingkaran adalah %.1f\n", areaMain)
	fmt.Printf("Keliling lingkaran adalah %.1f", kelilingMain)
}

func luasMultipleReturn(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d
	return area, circumference
}

func predefindeReturn(d float64) (area float64, circumferece float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumferece = math.Pi * d
	return
}
