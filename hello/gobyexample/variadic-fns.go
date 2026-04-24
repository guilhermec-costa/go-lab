package gobyexample

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums)

	for _i, n := range nums {
		fmt.Printf("Argument %v: %v\n", _i, n)
	}
}

func VariadicFunctions() {
	sum(1, 3, 4)
	fmt.Println("==========")

	scores := []int{1, 2, 3, 4}

	scoresSlice := make([]int, 0, 4)

	sum(scores...)
	fmt.Println(scoresSlice)
}
