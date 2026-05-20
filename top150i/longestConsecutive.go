package top150i

func longestConsecutive(nums []int) int {

	numsMap := make(map[int]bool)

	// set all true
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = true
	}

	// mark false which is not a starting point
	for v := range numsMap {
		if _, exist := numsMap[v-1]; exist {
			numsMap[v] = false
		}
	}

	// calculate seq
	ans := 0
	for v := range numsMap {
		if numsMap[v] {
			l := 1

			for _, ok := numsMap[v+l]; ok; _, ok = numsMap[v+l] {
				l++
			}

			if l > ans {
				ans = l
			}
		}
	}

	return ans

}
