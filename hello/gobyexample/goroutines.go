package gobyexample

import (
	"fmt"
	"time"
)

func fromStr(from string, sleep bool) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	fromStr("direct", false)

	// lightweight thread that CPU can apply concurrency
	go fromStr("goroutine", true)

	num := 5
	go func(msg string) {
		num += 3
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)

	fmt.Println("Value of num: ", num)
	fmt.Println("done")
}
