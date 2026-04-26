package gobyexample

import (
	"fmt"
	"time"
)

func Channels() {
	msgCh := make(chan string)

	go func() {
		msgCh <- "ping" // send "ping" into the channel
	}()

	time.Sleep(time.Second)
	msg := <-msgCh // receive message "ping" from the channel. This is blocking by default
	fmt.Println("Message from the channel ", msg)

	// buffered channel. Can send up 2 messages without having a corresponding concurrent receive
	messages := make(chan string, 2)

	messages <- "message 1"
	messages <- "messages 2"
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
}
