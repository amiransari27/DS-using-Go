package dp

func LengthOfLIS(nums []int) int {

	t := make([][]int, len(nums)+1)
	for i := range t {
		t[i] = make([]int, len(nums)+1)
		for j := range t[i] {
			t[i][j] = -1
		}
	}

	return solveLengthOfLIS(nums, 0, -1, t)
}

func solveLengthOfLIS(nums []int, i int, pi int, t [][]int) int {

	if i >= len(nums) {
		return 0
	}

	if pi != -1 && t[i][pi] != -1 {
		return t[i][pi]
	}

	var take int

	if pi == -1 || nums[i] > nums[pi] {
		take = 1 + solveLengthOfLIS(nums, i+1, i, t)
	}
	skip := solveLengthOfLIS(nums, i+1, pi, t)

	if pi != -1 {
		t[i][pi] = max(take, skip)
	}

	return max(take, skip)
}

func LengthOfLISBU(nums []int) int {

	n := len(nums)

	if n == 0 {
		return 0
	}
	t := make([]int, n)
	for k := range t {
		t[k] = 1
	}

	result := 1

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {

				t[i] = max(t[i], t[j]+1)

				result = max(result, t[i])
			}
		}
	}

	return result
}
