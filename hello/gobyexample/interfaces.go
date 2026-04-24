package gobyexample

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with radius ", c)
	}
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := square{width: 20, height: 20}
	measure(r)

	c := circle{radius: 20}
	measure(c)

	detectCircle(r);
	detectCircle(c);
}
