package leetcode

func RomanToInt(s string) int {
	ans, i := 0, 0

	for i < len(s) {
		switch s[i] {
		case 'M':
			ans += 1000
		case 'D':
			ans += 500
		case 'C':
			if i+1 < len(s) && (s[i+1] == 'M' || s[i+1] == 'D') {
				ans -= 100
			} else {
				ans += 100
			}

		case 'L':
			ans += 50
		case 'X':
			if i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
				ans -= 10
			} else {
				ans += 10
			}
		case 'V':
			ans += 5
		case 'I':
			if i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
				ans -= 1
			} else {
				ans += 1
			}
		}
		i++
	}

	return ans
}

func RomanToInt2(s string) int {
	ans, i := 0, 0

	for i < len(s) {
		if s[i] == "M"[0] {
			ans += 1000
		} else if s[i] == "D"[0] {
			ans += 500
		} else if s[i] == "C"[0] && i+1 < len(s) && (s[i+1] == "M"[0] || s[i+1] == "D"[0]) {
			ans -= 100
		} else if s[i] == "C"[0] {
			ans += 100
		} else if s[i] == "L"[0] {
			ans += 50
		} else if s[i] == "X"[0] && i+1 < len(s) && (s[i+1] == "L"[0] || s[i+1] == "C"[0]) {
			ans -= 10
		} else if s[i] == "X"[0] {
			ans += 10
		} else if s[i] == "V"[0] {
			ans += 5
		} else if s[i] == "I"[0] && i+1 < len(s) && (s[i+1] == "V"[0] || s[i+1] == "X"[0]) {
			ans -= 1
		} else if s[i] == "I"[0] {
			ans += 1
		}
		i++
	}

	return ans
}
