package recusion

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var solve func()
	solve = func() {
		if len(temp) == len(nums) {
			result = append(result, append([]int{}, temp...))
			return
		}

		for i, v := range nums {
			if used[i] {
				continue
			}

			temp = append(temp, v)
			used[i] = true
			solve()
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}

	solve()
	return result
}
