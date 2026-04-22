package gobyexample;

import (
	"fmt" 
	"math/rand"
)

func getRandom() int {
	return rand.Int();
}

func IfElse() {
	PrintFuncHeader("IfElse")
	var num uint32;
	fmt.Print("Type any number: ");
	fmt.Scan(&num)

	fmt.Println("Number typed: ", num);
	if(num % 2 == 0) {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd");
	}

	// n is only available at the condition scope. Avoid polluting global scope 
	if n := getRandom(); n % 2 == 0 {
		fmt.Println("Random number ", n, " is even");
	} else {
		fmt.Println("Random number ", n, " is odd");
	}
}