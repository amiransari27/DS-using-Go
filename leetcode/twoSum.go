package leetcode

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

type Pair struct {
	index int
	val   int
}

func TwoSumOptimized(nums []int, target int) []int {

	var pairArr = make([]Pair, 0)
	for i := 0; i < len(nums); i++ {
		pairArr = append(pairArr, Pair{i, nums[i]})
	}

	pairArr = quickSort(pairArr)

	low := 0
	high := len(nums) - 1

	for low < high {
		if pairArr[low].val+pairArr[high].val == target {
			return []int{pairArr[low].index, pairArr[high].index}
		} else if pairArr[low].val+pairArr[high].val < target {
			low++
		} else {
			high--
		}
	}

	return []int{}
}

func quickSort(pairArr []Pair) []Pair {

	sort(pairArr, 0, len(pairArr)-1)
	return pairArr
}

func sort(pairArr []Pair, low int, high int) {
	if low < high {
		pi := partision(pairArr, low, high)

		sort(pairArr, low, pi-1)
		sort(pairArr, pi+1, high)
	}
}

func partision(pairArr []Pair, low int, high int) int {

	pivot := pairArr[high]
	position := low - 1

	for i := low; i < high-1; i++ {
		if pairArr[i].val <= pivot.val {
			position++

			tmp := pairArr[position]
			pairArr[position] = pairArr[i]
			pairArr[i] = tmp
		}
	}

	position++
	tmp := pairArr[position]
	pairArr[position] = pairArr[high]
	pairArr[high] = tmp

	return position
}
