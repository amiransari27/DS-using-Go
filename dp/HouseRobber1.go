package dp

func Rob(nums []int) int {

	t := make([]int, len(nums))
	for i := range t {
		t[i] = -1
	}
	return solveRob(nums, 0, t)
}

func solveRob(nums []int, i int, t []int) int {

	if i >= len(nums) {
		return 0
	}
	if t[i] != -1 {
		return t[i]
	}

	steal := nums[i] + solveRob(nums, i+2, t)
	leave := solveRob(nums, i+1, t)

	t[i] = max(steal, leave)

	return t[i]

}

func solveBU(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	t := make([]int, n)

	t[0] = nums[0]
	t[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		t[i] = max(nums[i]+t[i-2], t[i-1])
	}

	return t[n-1]
}
