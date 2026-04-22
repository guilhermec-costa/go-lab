package gobyexample

import (
	"fmt"
)

type Myint int

func Slices() {
	PrintFuncHeader("Slices")

	// slice declaration
	var s []string
	// slice initialization with capacity 3. Allocates memory for the slice
	s = make([]string, 3)

	fmt.Println("Starting slice state: ", s, " len: ", len(s), " cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f", "g", "h")
	fmt.Println("set :", s)

	options := make([]string, 5)

	options[0] = "hello"
	for _opIdx := range len(options) {
		fmt.Println("Option: ", options[_opIdx])
	}

	// no capacity allocatd. Append will handle this at the first time
	var temperatures []float32
	temperatures = append(temperatures, 20.5, 30.0, 35.6, 41.2, 26.8)
	fmt.Println("temps: ", temperatures)

	cp_tmps := make([]float32, len(temperatures))
	copy(cp_tmps, temperatures)
	fmt.Println("temps copies: ", cp_tmps)
	fmt.Println("Sliced copies: ", cp_tmps[(len(cp_tmps)/2):])

	// slice declaration
	var cp_tmps2 []float32;
	// allocates capacity for the slice
	cp_tmps2 = make([]float32, len(cp_tmps));
	cp_tmps2 = cp_tmps[2:];
	fmt.Println("Sliced copies2: ", cp_tmps2)

	cp_tmps3 := cp_tmps2[:];
	fmt.Println("Sliced copies3: ", cp_tmps3)
}
