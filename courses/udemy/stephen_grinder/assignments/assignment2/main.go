package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func main() {
	tr := triangle{
		base:   2.0,
		height: 4.0,
	}
	fmt.Println("Triangle area", tr.getArea())
	s := square{
		sideLength: 5.00,
	}
	fmt.Println("Square area", s.getArea())
}
