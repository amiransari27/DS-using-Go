package recusion

func permuteUnique(nums []int) [][]int {

	result := make([][]int, 0)
	numMap := make(map[int]int)
	for _, n := range nums {
		numMap[n] += 1
	}
	uniqueNums := make([]int, 0, len(numMap))
	for n := range numMap {
		uniqueNums = append(uniqueNums, n)
	}
	tmp := make([]int, 0, len(nums))

	var solve func()

	solve = func() {
		if len(tmp) == len(nums) {
			result = append(result, append([]int{}, tmp...))
			return
		}

		for _, v := range uniqueNums {
			if o := numMap[v]; o > 0 {
				tmp = append(tmp, v)
				numMap[v] -= 1
				solve()
				numMap[v] += 1
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	solve()

	return result
}

func permuteUniqueSwap(nums []int) [][]int {

	result := make([][]int, 0)

	var solve func(idx int)

	solve = func(idx int) {

		if idx == len(nums) {
			result = append(result, append([]int{}, nums...))
			return
		}

		used := make(map[int]bool)
		for i := idx; i < len(nums); i++ {
			candidate := nums[i]
			if used[candidate] {
				continue
			}
			used[candidate] = true
			nums[i], nums[idx] = nums[idx], nums[i]
			solve(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}

	}

	solve(0)

	return result
}
