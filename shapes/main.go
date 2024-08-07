package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sidelength float64
}

type triangle struct {
	length float64
	height float64
}

func main() {
	t := triangle{length: 5, height: 7}
	h := square{sidelength: 10}

	printArea(t)
	printArea(h)
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.length
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
