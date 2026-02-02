package dp

func MaxAlternatingSum(nums []int) int64 {

	n := len(nums)
	t := make([][]int64, n)
	for i := range t {
		t[i] = make([]int64, 2)
		for j := range t[i] {
			t[i][j] = -1
		}
	}

	return solveMaxAlternatingSum(nums, 0, true, t)
}

func solveMaxAlternatingSum(nums []int, i int, flag bool, t [][]int64) int64 {

	if i >= len(nums) {
		return 0
	}

	flag_j := map[bool]int{true: 1, false: 0}[flag]
	if t[i][flag_j] != -1 {
		return t[i][flag_j]
	}

	skip := solveMaxAlternatingSum(nums, i+1, flag, t)

	val := nums[i]
	if flag == false {
		val = -val
	}
	take := solveMaxAlternatingSum(nums, i+1, !flag, t) + int64(val)

	t[i][flag_j] = max(skip, take)

	return t[i][flag_j]
}

func MaxAlternatingSumBU(nums []int) int64 {

	n := len(nums)
	t := make([][2]int64, n+1)

	for i := 1; i < n+1; i++ {

		t[i][0] = max(t[i-1][1]-int64(nums[i-1]), t[i-1][0])
		t[i][1] = max(t[i-1][0]+int64(nums[i-1]), t[i-1][1])
	}

	return max(t[n][0], t[n][1])
}

func MaxAlternatingSumBUConstSpace(nums []int) int64 {

	var even, odd int64

	for i := 0; i < len(nums); i++ {

		even = max(odd-int64(nums[i]), even)
		odd = max(even+int64(nums[i]), odd)
	}

	return max(even, odd)
}
