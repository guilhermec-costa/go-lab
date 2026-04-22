package main

import (
	"fmt"
	"gofromscratch/gobyexample"
	"log"
	"example.com/greetings"
	"rsc.io/quote"
)

func expl() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0);

	fmt.Println("Starting example program...");
	msg, err := greetings.Hello("Churros");
	if(err != nil) {
		log.Fatal(err);
	}

	fmt.Println(msg);
	fmt.Printf("Line from go filosofy: %v\n", quote.Go());

	_, err2 := greetings.Hello("");
	if(err2 != nil) {
		log.Fatal(err2);
	}

	fmt.Println("End of program");
}

func main()  {
	gobyexample.Main();
}