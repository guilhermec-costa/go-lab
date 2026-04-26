package gobyexample

import (
	"fmt"
	"time"
)

func Tickers() {
	ticker1 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return

			case c := <-ticker1.C:
				fmt.Println("tick at", c)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	done <- true
	fmt.Println("ticker stopped")
}
