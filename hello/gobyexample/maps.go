package gobyexample

import (
	"fmt"
)

func Maps() {

	students := make(map[string]uint32)
	students["Guilherme"] = 21
	students["Isabela"] = 12

	dogs := map[string]uint32{
		"Churros": 12,
		"Shoyou":  11,
	}

	fmt.Println("Students map: ", students)
	fmt.Println("Dogs map: ", dogs)

	delete(students, "Guilherme")
	fmt.Println("Students map after deletion: ", students)

	clear(dogs)
	fmt.Println("Dogs map after clear: ", dogs)
}
