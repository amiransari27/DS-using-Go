package leetcode75

func SingleNumber(nums []int) int {

	xor := 0

	for _, n := range nums {
		xor = xor ^ n
	}

	return xor
}
