package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {

	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// if the number is not prime you will find a divisor at less than or equal of that number
func isPrimeOptimize(n int) bool {

	if n < 2 {
		return false
	}

	for i := 2.0; i <= math.Sqrt(float64(n)); i++ {
		fmt.Println(float64(n), i)
		if math.Mod(float64(n), i) == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPrime(1)) // false
	fmt.Println(isPrime(2)) // true
	fmt.Println(isPrime(4)) // fasle
	fmt.Println(isPrime(5)) // true

	fmt.Println("--------")

	fmt.Println(isPrimeOptimize(1)) // false
	fmt.Println(isPrimeOptimize(2)) // true
	fmt.Println(isPrimeOptimize(4)) // fasle
	fmt.Println(isPrimeOptimize(5)) // true

}
