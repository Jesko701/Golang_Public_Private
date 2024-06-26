package main

import (
	"fmt"
	"math"
)

type calculate interface {
	broad() float64
	circumferential() float64
	// ! no radius for ⭕
}

// Circle struct
type circle struct {
	diameter float64
}

func (c circle) radius() float64 {
	return c.diameter / 2 // * function will be error if it's called inside interface
}

func (c circle) broad() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) circumferential() float64 {
	return math.Pi * c.diameter
}

// Rectangle struct
type rectangle struct {
	square_side float64
}

func (r rectangle) broad() float64 {
	return math.Pow(r.square_side, 2)
}

func (r rectangle) circumferential() float64 {
	return r.square_side * 4
}

func main() {
	var calculate_plane calculate

	// Circle
	calculate_plane = circle{14.0}
	fmt.Println("===== Circle")
	fmt.Println("Broad      	:", calculate_plane.broad())
	fmt.Println("Circumferece  	:", calculate_plane.circumferential())
	// fmt.Println("Radius :", calculate_plane.radius()) // Error: calculate interface has no method 'radius'
	// * this how to handle function that is none in interface
	if c, ok := calculate_plane.(circle); ok {
		fmt.Println("Radius		:", c.radius())
	}

	// Rectangle
	calculate_plane = rectangle{4}
	fmt.Println("===== Rectangle")
	fmt.Println("Broad      	:", calculate_plane.broad())
	fmt.Println("Circumference  	:", calculate_plane.circumferential())
}
