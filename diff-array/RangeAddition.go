package diffarray

/**
 * @param length: the length of the array
 * @param updates: update operations
 * @return: the modified array after all k operations were executed
 */
func GetModifiedArray(length int, updates [][]int) []int {
	// Write your code here
	result := make([]int, length)

	for _, values := range updates {
		s, e, x := values[0], values[1], values[2]
		result[s] += x

		if e+1 < length {
			result[e+1] -= x
		}
	}

	for i := 1; i < length; i++ {
		result[i] = result[i-1] + result[i]
	}

	return result
}
