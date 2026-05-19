package top150i

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1

	res := make([]int, 2)
	for l < r {
		if numbers[l]+numbers[r] == target {
			res[0] = l + 1
			res[1] = r + 1
			break
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}

	}

	return res
}

func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		re := target - nums[i]

		if j, exist := numsMap[re]; exist {
			return []int{j, i}
		} else {
			numsMap[nums[i]] = i
		}
	}

	return []int{}
}
