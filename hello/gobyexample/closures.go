package gobyexample

import (
	"fmt"
)

func counter() func() int {
	c := 0;
	return func() int {
		c++;
		return c;
	}
}

func Closures() {
	nextInt := counter();

	fmt.Println(nextInt());
	fmt.Println(nextInt());
	fmt.Println(nextInt());
	fmt.Println(nextInt());

	nextInt2 := counter();
	fmt.Println(nextInt2());
	fmt.Println(nextInt2());
}