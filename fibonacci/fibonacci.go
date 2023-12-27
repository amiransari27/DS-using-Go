package main

import "fmt"

func fibonacci(n int) []int {
	fib := []int{0, 1}

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib
}

func main() {
	fmt.Println("main")
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(7))

}
