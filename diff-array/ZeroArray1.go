package diffarray

func ZeroArray1(nums []int, queries [][]int) bool {
	n := len(nums)

	diff := make([]int, n)

	for _, vals := range queries {
		l, r, x := vals[0], vals[1], 1

		diff[l] += x
		if r+1 < n {
			diff[r+1] -= x
		}
	}

	diffSum := 0
	for i, v := range diff {
		diffSum += v
		diff[i] = diffSum
	}

	for i, v := range diff {
		if v < nums[i] {
			return false
		}
	}
	return true
}
