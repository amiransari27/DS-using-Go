package recusion

func permute(nums []int) [][]int {

	result := make([][]int, 0)

	var solve func(temp *[]int, st *map[int]bool)

	solve = func(temp *[]int, st *map[int]bool) {

		if len(*temp) == len(nums) {
			result = append(result, append([]int{}, *temp...))
			return
		}

		for _, v := range nums {

			if isTrue := (*st)[v]; !isTrue {

				(*temp) = append((*temp), v)
				(*st)[v] = true
				solve(temp, st)

				(*temp) = (*temp)[:len(*temp)-1]
				(*st)[v] = false

			}

		}
	}

	temp := make([]int, 0)
	st := make(map[int]bool)
	solve(&temp, &st)

	return result
}
