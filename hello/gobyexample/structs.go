package gobyexample

import (
	"fmt"
)

type person struct {
	name string
	age  uint8
}

// idiomatic
func newPerson(name string) *person {
	return &person{
		name: name,
		age:  42,
	}
}

func Structs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Churros", age: 12})
	user := &person{"Guilherme", 20}
	fmt.Println(user)

	fmt.Println(newPerson("Shoyou"))

	dog := struct {
		name string;
		cool bool
	} {
		name: "Churros",
		cool: true,
	}

	fmt.Println("Dog: ", dog);
}
