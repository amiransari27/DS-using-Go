package leetcode

import (
	"math"
	"strings"
)

func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var ans float64
	isNeg := false
	for i := 0; i < len(s); i++ {
		charCode := s[i]
		if i == 0 {
			if s[i] == "-"[0] {
				isNeg = true
				continue
			} else if s[i] == "+"[0] {
				isNeg = false
				continue
			}
		}

		if charCode >= "0"[0] && charCode <= "9"[0] {

			digit := charCode - "0"[0]

			ans = ans*10 + float64(digit)

		} else {
			break
		}
	}

	if isNeg {
		ans = -ans
	}

	if ans < math.MinInt32 {
		return math.MinInt32
	} else if ans > math.MaxInt32 {
		return math.MaxInt32
	}

	return int(ans)
}
