package trickque

import (
	"fmt"
	"sync"
)

func FanOutAndIn() {
	jobs := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	res1 := work(jobs)
	res2 := work(jobs)
	res3 := work(jobs)

	for res := range merge(res1, res2, res3) {
		fmt.Println("result:", res)
	}

}

func work(jobs <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for j := range jobs {
			out <- j * j
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {

	merged := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range c {
				merged <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
