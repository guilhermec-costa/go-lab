package gobyexample

import (
	"fmt"
)

type base struct {
	num int
}

type container struct {
	base // struct embedding
	str  string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func StructEmbeddings() {

	type describer interface {
		describe() string
	}

	co := container{
		// struct embedding
		base: base{
			num: 2,
		},
		str: "Some str",
	}

	fmt.Println("Describe from co: ", co.describe())
	var de describer = co
	fmt.Println("Describer from de: ", de.describe()) // same thing, but scoped to the describer interface
}
