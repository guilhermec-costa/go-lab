package gobyexample

import (
	"fmt"
	"time"
)

func _worker(jobId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", jobId, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", jobId, "finished job", j)
		results <- j * 2
	}
}

func WorkerPools() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go _worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
