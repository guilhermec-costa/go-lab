package gobyexample

import (
	"fmt"
)

func NonBlocking() {
	messages := make(chan string)
	// signals := make(chan string)

	select {
	// nonblocking receive. Even the channel being unbuffered, default case will be selected
	// if there's no message on messages channel
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	// non blocking send. Even the channel being unbuffered,
	// default cause will be selected if there's no correponding receiver
	case messages <- msg:
		fmt.Println("sent message", msg)

	default:
		fmt.Println("no message sent")
	}

}
