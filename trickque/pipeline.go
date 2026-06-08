package trickque

import "fmt"

func generate(nums ...int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()
	return out
}

func square(ch chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range ch {
			out <- v * v
		}
	}()
	return out
}

func filterEven(ch chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range ch {
			if v%2 == 0 {
				out <- v
			}
		}
	}()
	return out
}

func Pipeline() {

	nums := generate(1, 2, 3, 4, 5, 6)
	squared := square(nums)
	filtered := filterEven(squared)

	for v := range filtered {
		fmt.Println("result ", v)
	}

}
