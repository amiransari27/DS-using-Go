package problems

func IsPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}

	return true
}

// if n is power of two
// when we run a bit and operation between n&(n-1) it always return 0
func IsPowerOfTwoBitWise(n int) bool {
	if n < 1 {
		return false
	}

	return (n & (n - 1)) == 0
}
