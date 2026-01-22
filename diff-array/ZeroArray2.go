package diffarray

func ZeroArray(nums []int, queries [][]int) int {
	//find
	if allZero(nums) {
		return 0
	}

	q := len(queries)

	for i := 0; i < q; i++ {
		if checkWithDiffArrayTech(nums, queries, i) {
			return i + 1
		}
	}
	return -1
}

func checkWithDiffArrayTech(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	diff := make([]int, n+1)

	for i := 0; i <= k; i++ {
		l, r, x := queries[i][0], queries[i][1], queries[i][2]

		diff[l] += x
		if r+1 <= n {
			diff[r+1] -= x
		}
	}

	sum := 0
	for j := 0; j < n; j++ {
		sum += diff[j]
		diff[j] = sum

		if nums[j]-diff[j] > 0 {
			return false
		}
	}

	return true

}

func allZero(nums []int) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}

	return true
}
