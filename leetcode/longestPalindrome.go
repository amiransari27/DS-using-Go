package leetcode

func LongestPalindrome(s string) string {
	//abba
	//0123
	maxCount := 0
	maxStr := ""
	for i := 0; i < len(s); i++ {
		tmpCount := 0
		for j := i; j <= len(s); j++ {
			tmpCount++
			str := s[i:j]
			if isPalindrome(str) {
				if maxCount < tmpCount {
					maxCount = tmpCount
					maxStr = str
				}
			}
		}
	}
	return maxStr
}

func LongestPalindromeOptimized(s string) string {
	result := ""
	//odd orbit
	for ax := 0; ax < len(s); ax++ {
		orbit := 1
		length := 1

		for ax-orbit >= 0 && ax+orbit < len(s) {
			if s[ax-orbit] == s[ax+orbit] {
				orbit++
				length += 2
			} else {
				break
			}
		}

		if len(result) < length {
			low := ax - length/2
			high := ax + length/2 + 1
			result = s[low:high]
		}

	}

	//even orbit

	for ax := 0; ax < len(s)-1; ax++ {
		orbit := 1
		length := 0

		for ax-orbit+1 >= 0 && ax+orbit < len(s) {
			if s[ax-orbit+1] == s[ax+orbit] {
				orbit++
				length += 2
			} else {
				break
			}
		}

		if len(result) < length {
			low := ax - length/2 + 1
			high := ax + length/2 + 1
			result = s[low:high]
		}

	}

	return result
}

func isPalindrome(s string) bool {
	low := 0
	high := len(s) - 1

	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}
