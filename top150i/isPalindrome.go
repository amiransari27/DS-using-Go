package top150i

import "strings"

func isPalindrome(s string) bool {

	s = strings.ToLower(s)

	l, r := 0, len(s)-1

	for l < r {
		l_char := s[l]
		r_char := s[r]

		if !((l_char >= 'a' && l_char <= 'z') || (l_char >= '0' && l_char <= '9')) {
			l++
			continue
		}

		if !((r_char >= 'a' && r_char <= 'z') || (r_char >= '0' && r_char <= '9')) {
			r--
			continue
		}

		if l_char != r_char {
			return false
		}

		l++
		r--

	}

	return true
}
