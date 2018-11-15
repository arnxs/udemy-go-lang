package main

import "fmt"

type shape interface {
	getArea() float64
	printArea()
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	s := square{sideLength: 10}

	s.printArea()
	printArea(s)

	t := triangle{base: 10, height: 10}

	t.printArea()
	printArea(t)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) printArea() {
	fmt.Printf("Area of triangle base=%v height=%v is %v\n", t.base, t.height, t.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) printArea() {
	fmt.Printf("Area of square side=%v is %v\n", s.sideLength, s.getArea())
}

func printArea(s shape) {
	s.printArea()
}
