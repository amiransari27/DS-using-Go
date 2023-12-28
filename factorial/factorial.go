import "fmt"

func factorial(n int) int {
	var fact int = 1

	for n >= 1 {
		fact = fact * n
		n--
	}
	return fact
}

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(2))
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
}
