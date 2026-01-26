package dp

func Fib(n int) int {

	t := make([]int, n+1)
	for i := range t {
		t[i] = -1
	}
	return solveFib(n, t)

}

func solveFib(n int, t []int) int {
	if n <= 1 {
		return n
	}
	if t[n] != -1 {
		return t[n]
	}
	t[n] = solveFib(n-1, t) + solveFib(n-2, t)
	return t[n]
}

func FibBottomUp(n int) int {
	if n <= 1 {
		return n
	}
	t := make([]int, n+1)
	t[0] = 0
	t[1] = 1

	for i := 2; i <= n; i++ {
		t[i] = t[i-1] + t[i-2]
	}

	return t[n]

}
