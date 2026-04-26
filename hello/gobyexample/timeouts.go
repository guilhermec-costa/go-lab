package gobyexample

import (
	"fmt"
	"time"
)

func Timeouts() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// one-time select
	select {
	case res := <-c1:
		fmt.Println("res: ", res)

	// timeout
	case curTime := <-time.After(1 * time.Second):
		fmt.Println("timeout 1", curTime)
	}
}
