package top150i

import "fmt"

func summaryRanges(nums []int) []string {

	res := make([]string, 0)

	for i := 0; i < len(nums); {

		j := i + 1

		for j < len(nums) && nums[j-1]+1 == nums[j] {
			j++
		}

		if i == j-1 {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
		}

		i = j
	}

	return res
}
