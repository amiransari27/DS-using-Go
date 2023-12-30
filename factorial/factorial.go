package main

import "fmt"

func factorial(n int) int {
	var fact int = 1

	for n >= 1 {
		fact = fact * n
		n--
	}
	return fact
}

func factorialRec(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRec(n-1)
}

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(2))
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println("-----")

	fmt.Println(factorialRec(0))
	fmt.Println(factorialRec(2))
	fmt.Println(factorialRec(4))
	fmt.Println(factorialRec(5))
}
