package leetcode

import (
	"fmt"
)

//O(n2)
func LengthOfLongestSubstring(s string) int {
	var maxCount int = 0
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
	var obj map[string]bool

	for i := 0; i < len(s); i++ {
		obj = make(map[string]bool)
		tmpCount := 0
		for j := i; j < len(s); j++ {

			if !obj[string(s[j])] {
				obj[string(s[j])] = true
				tmpCount++
			} else {
				break
			}
		}

		if maxCount < tmpCount {
			maxCount = tmpCount
		}
	}
	return maxCount
}

// O(2n) - O(n)
func LengthOfLongestSubstringOptimized(s string) int {
	var maxCount int = 0
	// when we try to read string char using index it returns ascci vallue as uint8
	var obj = make(map[uint8]bool)
	var j int = -1
	for i := 0; i < len(s); i++ {

		if !obj[s[i]] {
			obj[s[i]] = true
		} else {
			for {
				j++
				delete(obj, s[j])
				if s[j] == s[i] {
					i--
					break
				}
			}
		}
		tmpCount := i - j
		maxCount = max(maxCount, tmpCount)
	}

	return maxCount
}
