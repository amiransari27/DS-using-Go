package trickque

import (
	"fmt"
	"sync"
)

func worker(workerId int, jobs chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result := job * job
		fmt.Printf("Worker %d is processing job %d \n", workerId, job)
		results <- result
	}

}

func WorkerPool() {

	numWorker := 3
	numJobjs := 100

	jobs := make(chan int, numJobjs)
	results := make(chan int, numJobjs)
	var wg sync.WaitGroup

	// run all workers
	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// send all jobs

	for i := 1; i <= numJobjs; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}

}
