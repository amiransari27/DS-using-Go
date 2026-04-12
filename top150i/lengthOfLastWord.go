package top150i

import "strings"

func lengthOfLastWord(s string) int {

	s = strings.Trim(s, " ")

	n := len(s)
	count := 0
	for i := n - 1; i >= 0; i-- {

		if s[i] == ' ' {
			break
		}
		count++
	}

	return count
}
