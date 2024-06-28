package main

import "math"

type cube struct {
	side float64
}

func (c cube) broad() float64 {
	return math.Pow(c.side, 2) * 6
}

func (c cube) volume() float64 {
	return math.Pow(c.side, 3)
}

func (c cube) circumference() float64 {
	return c.side * 12
}

func main() {

}
