package gobyexample;

import (
	"fmt"
	"math"
)

const s string = "constant";

func Constants() {
	PrintFuncHeader("Constants");

	fmt.Println("Constant value: ", s);
	const flags uint32 = 0b00000101;
	const hex uint8 =  0x000000FF;

	const d = 500;
	fmt.Println("Flags: ", flags);	
	fmt.Println("Flags as uint64: ", uint64(flags));	
	fmt.Println("Flags as hex: ", hex);	

	const ang float64 = 90;
	sin := math.Sin(ang);

	fmt.Println("const number: ", d);
	fmt.Println("Sin of angle ", ang, " = ", math.Ceil(sin));
}