package top150

import "strings"

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		ch1 := s[l]
		ch2 := s[r]

		if !((ch1 >= 'a' && ch1 <= 'z') || (ch1 >= '0' && ch1 <= '9')) {
			l++
			continue
		}
		if !((ch2 >= 'a' && ch2 <= 'z') || (ch2 >= '0' && ch2 <= '9')) {
			r--
			continue
		}

		if ch1 != ch2 {
			return false
		}

		l++
		r--
	}

	return true
}
