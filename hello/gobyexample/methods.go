package gobyexample

import (
	"fmt"
)

type rect struct {
	width, height uint32
}

func (r *rect) area() uint32 {
	return r.width * r.height
}

func (r rect) perim() uint32 {
	return 2*r.width + 2*r.width
}

func Methods() {
	terrain := rect{width: 20, height: 20}

	fmt.Printf("Area of %v: %v\n", terrain, terrain.area())
	fmt.Printf("Perimeter of %v: %v\n", terrain, terrain.perim())

	rp := &terrain

	fmt.Printf("Area of %v: %v\n", rp, rp.area())
	fmt.Printf("Perimeter of %v: %v\n", rp, rp.perim())

	fmt.Printf("Address of terrain: %x\n", &terrain)
	fmt.Printf("Address of rp: %x\n", rp)
}
