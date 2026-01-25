package diffarray

import "slices"

func MaxFrequency(nums []int, k int, numOperations int) int {

	maxVal := slices.Max(nums)

	diff := make([]int, maxVal+2)
	fq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		fq[nums[i]] += 1

		l := max(nums[i]-k, 0)
		r := min(nums[i]+k, maxVal)
		diff[l]++
		diff[r+1]--
	}

	result := 1

	for target := 0; target <= maxVal; target++ {
		if target > 0 {
			diff[target] += diff[target-1]
		}

		currFq := fq[target]
		needConversion := diff[target] - currFq
		maxPosible := min(needConversion, numOperations)

		result = max(result, currFq+maxPosible)
	}

	return result
}
