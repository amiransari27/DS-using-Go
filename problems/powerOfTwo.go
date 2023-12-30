package problems

import "math"

func IsPowerOfTwo(n float64) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if math.Mod(n, 2) != 0 {
			return false
		}
		n = n / 2
	}

	return true
}
