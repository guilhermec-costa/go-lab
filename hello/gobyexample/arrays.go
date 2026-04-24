package gobyexample

import (
	"fmt"
)

func Arrays() {
	PrintFuncHeader("Arrays")
	var a [5]int
	fmt.Println("Array: ", a)

	a[4] = 100
	fmt.Println("Array: ", a)

	b := [3]int{1, 2, 3}
	fmt.Println("Array b: ", b)

	c := [...]any{"i", "built", "this", "program", 5, false}
	for i := range(len(c)) {
		fmt.Print(c[i], " ");
	}
	fmt.Println();

	// slice. Dynamic size
	var d []string;
	fmt.Println("Slice d: ", d == nil, len(d) == 0);

	tmps := [...]float32{30.5, 20, 22.2,26.8};
	fmt.Println("Sliced array: ", tmps[2:])

	for _, t := range tmps {
		fmt.Printf("Temperature %v\n", t);
	}
}
