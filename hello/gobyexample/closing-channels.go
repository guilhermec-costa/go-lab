package gobyexample

import "fmt"

func ClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan struct{})

	go func() {
		for {
			// j is the int value
			// the more value will be false if jobs has been closed and all values in the channel have already been received
			if j, more := <-jobs; more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- struct{}{}
				return // ends the goroutine
			}
		}
	}()

	// equivalent version
	// go func() {
	// 	for v := range jobs {
	// 		fmt.Println("received job", v)
	// 	}
	// 	fmt.Println("done reading")
	// 	done <- struct{}{}
	// }()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// closing a channel indicates that no more values will be sent on it
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
