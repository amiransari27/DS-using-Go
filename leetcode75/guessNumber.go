package leetcode75

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func GuessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := (low + high) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func guess(num int) int {
	return 0
}
