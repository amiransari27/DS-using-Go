package diffarray

import "slices"

func MaxFrequency2(nums []int, k int, numOperations int) int {

	maxVal := slices.Max(nums)
	diffMap := make(map[int]int)
	fq := make(map[int]int)
	for _, num := range nums {
		fq[num] += 1

		l := max(num-k, 0)
		r := min(num+k, maxVal)

		diffMap[l] += 1
		diffMap[num] += 0
		diffMap[r+1] -= 1

	}

	i := 0
	order := make([]int, len(diffMap))
	for key := range diffMap {
		order[i] = key
		i++
	}
	slices.Sort(order)

	result := 1

	var pK int
	for j, key := range order {
		if j > 0 {
			diffMap[key] = diffMap[pK] + diffMap[key]
		}

		currF := fq[key]
		needF := diffMap[key] - currF
		maxF := min(needF, numOperations)

		result = max(result, currF+maxF)

		pK = key
	}

	return result
}
