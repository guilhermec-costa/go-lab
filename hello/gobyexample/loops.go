package gobyexample

import (
	"fmt"
)

func Loops() {
	PrintFuncHeader("Loops");

	var i uint8 = 1;
	for i <= 3 {
		fmt.Println(i);
		i++;
	}

	for n:=0;n<10; {
		fmt.Println("n range: ", n);
		n++;
	}

	for j := 0; j < 3; j++ {
		fmt.Println("range", j)	;	
	}

	for z := range 3 {
		fmt.Println("range z", z);
	}

	for {
		fmt.Println("Infinite loop and then break");
		break;
	}

	fmt.Print("Type a number: ");

	var _i uint8;
	_, err := fmt.Scan(&_i);
	if(err != nil) {
		fmt.Println("Failed to read number", err);
	}

	fmt.Println("Number typed: ", _i);

	for x := range _i {
		fmt.Println("X value: ", x);
	}

	for x := 0; x<15;x++ {
		fmt.Println("X value: ", x);
	}
}