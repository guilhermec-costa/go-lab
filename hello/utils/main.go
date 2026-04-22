package utils

import "fmt"

func main() {
	fmt.Println("Hello from utils main");
}

func ExportedFunc() {
	fmt.Println("This function is exported");
}

func ExportedFunc2() {
	fmt.Println("This function is exported");
}