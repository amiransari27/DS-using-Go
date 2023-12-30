package main

import "fmt"

// time complexity O(n)
func fibonacci(n int) []int {
	fib := []int{0, 1}

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib
}

// time complexity - O(n)
// because at each level we are calling the function once
func fibonacciRec(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	fib := fibonacciRec(n - 1)
	fib = append(fib, fib[n-2]+fib[n-3])
	return fib
}

// return nth element of fibonacci series
// time complexity - O(2^n)
// because at each level we are calling the function twice
func fibonacciNthNumRec(n int) int {
	if n < 2 {
		return n
	}
	num := fibonacciNthNumRec(n-1) + fibonacciNthNumRec(n-2)
	return num
}

func main() {
	fmt.Println("main")
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(7))
	fmt.Println("-------")
	fmt.Println(fibonacciRec(2))
	fmt.Println(fibonacciRec(5))
	fmt.Println(fibonacciRec(7))
	fmt.Println("-------")
	fmt.Println(fibonacciNthNumRec(2))
	fmt.Println(fibonacciNthNumRec(5))
	fmt.Println(fibonacciNthNumRec(7))

}
