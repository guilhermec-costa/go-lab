package gobyexample

import (
	"fmt"
	"time"
)

func worker(done chan bool, dur time.Duration) {
	fmt.Print("Working...")
	time.Sleep(dur)
	fmt.Println("done")

	done <- true
	fmt.Println("Finished worker!")
}

func testWorker() {
	done := make(chan bool)

	go worker(done, time.Second)

	<-done
	<-done
	fmt.Println("finished testWorker fn")
}

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing", i)
		// time.Sleep(3 * time.Second)
		ch <- i
		fmt.Println("Sent", i)
	}
	close(ch)
}

func consumer(producer <-chan int, doneConsuming chan<- bool) {
	for v := range producer {
		time.Sleep(time.Second * 2)
		fmt.Println("Consumed", v)
	}

	doneConsuming <- true
}

func testBufferedChannel() {
	ch := make(chan int, 3)
	go producer(ch)

	consumed := make(chan bool)
	go consumer(ch, consumed)
	<-consumed

	fmt.Println("Done consuming producer values")
}

func ChannelSync() {
	// testWorker()
	//testBufferedChannel()

	prod := make(chan string, 1)

	go func() {
		prod <- "hello world"
		fmt.Println("Produced")
		time.Sleep(time.Second * 5)
		prod <- "hello world 2"
		fmt.Println("Produced")
		prod <- "hello world 3"
		fmt.Println("Produced")
		close(prod)
	}()

	for v := range prod {
		time.Sleep(time.Second * 3)
		fmt.Println("Value: ", v)
	}
}
