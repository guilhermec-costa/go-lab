package gobyexample

import (
	"fmt"
	"sync"
	"time"
)

func test1WaitGroups() {
	var wg sync.WaitGroup

	// adds 2 to delta
	wg.Add(2)

	go func() {
		fmt.Println("worker 1 working...")
		time.Sleep(time.Second * 4)
		wg.Done()
	}()

	go func() {
		fmt.Println("worker 2 working...")
		time.Sleep(time.Second * 2)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Finished workers")
}

func test2WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
		})
	}

	wg.Wait()
	fmt.Println("Finished workers")
}

func WaitGroups() {
	// test1WaitGroups()
	test2WaitGroups()
}
