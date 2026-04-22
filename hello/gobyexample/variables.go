package gobyexample

import "fmt"

func Variables() {
	fmt.Println("------ Variables ------");

	// declaration + initilization
	var message string = "hello world";
	var conditional bool = false;

	// multiple variables declaration/initialization
	var weight, age int = 60, 21;

	// go will infer the variable type
	var dogname = "churros";

	var amount int;
	amount = 15;
	// shorthand for declaration + initialization + type inference
	dogname2 := "shoyou";

	fmt.Println("Message: ", message);
	fmt.Println("Conditional: ", conditional);
	fmt.Println("Weight: ", weight);
	fmt.Println("Age: ", age);
	fmt.Println("Dog name: ", dogname);
	fmt.Println("Dog name2: ", dogname2);
	fmt.Println("Amount: ", amount);
}