package leetcode75

/*
right most bit (n & 1)
right shift -  (n >>=1)
*/
func MinFlips(a int, b int, c int) int {

	flip := 0
	for a != 0 || b != 0 || c != 0 {
		if c&1 == 0 {
			if a&1 == 1 {
				flip++
			}
			if b&1 == 1 {
				flip++
			}
		} else {
			if (a&1) == 0 && b&1 == 0 {
				flip++
			}
		}

		a >>= 1 // right shift by 1 bit
		// a = a >> 1
		b >>= 1
		c >>= 1
	}
	return flip
}
