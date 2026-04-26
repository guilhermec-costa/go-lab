package gobyexample

import (
	"fmt"
	"time"
)

func Selects() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// msg1 := <-c1;
	// msg2 := <-c2;

	// this would only be visible after both channels finished
	// fmt.Println("msg1", msg1)
	// fmt.Println("msg2", msg2)

	// channel multiplexer decoupled from order of message receiving
	// there's no "order". Receives values as they come
	for range 2 {
		select {
		case msg2 := <-c2:
			fmt.Println("received msg2", msg2)
			fmt.Println("done processing msg2", msg2)
		case msg1 := <-c1:
			fmt.Println("received msg1", msg1)
			fmt.Println("done processing msg1", msg1)
		}
	}
}
