package dp

func ClimbStairs(n int) int {

	t := make([]int, n+1)
	for i := range t {
		t[i] = -1
	}
	return solveClimbStairs(n, t)
}

func solveClimbStairs(n int, t []int) int {

	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if t[n] != -1 {
		return t[n]
	}

	t[n] = solveClimbStairs(n-1, t) + solveClimbStairs(n-2, t)

	return t[n]

}
