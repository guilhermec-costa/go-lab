package gobyexample

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C // read-only channel. It is notified after the timer fires
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	go func() {
		time.Sleep(5 * time.Second)
		// return a boolean indicating if the timer was stopped or not
		// if the timer was already fired, then it will not be stopped
		// if the timer is still pending, then it can be stopped
		stop2 := timer2.Stop()
		if stop2 {
			fmt.Println("timer 2 stopped")
		}
	}()

	time.Sleep(10 * time.Second)
}
