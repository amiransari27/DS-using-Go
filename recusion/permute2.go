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
