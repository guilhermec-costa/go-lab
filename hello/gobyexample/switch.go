package gobyexample

import (
	"fmt"
)

func Switch() {
	anon := func() {
		fmt.Println("Inside anonymous function")
	}
	anon();

	_getType := func(i any) {
		switch t := i.(type) {
		case bool:
			{
				fmt.Println("Boolean", t)
				break
			}
		case int:
			{
				fmt.Println("Integer", t)
				break
			}

		default:
			{
				fmt.Println("Unknown type", t)
			}
		}
	}

	_getType(true);
	_getType(1);
	_getType("hello world");
}
