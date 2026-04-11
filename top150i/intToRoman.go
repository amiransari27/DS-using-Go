package top150i

import "strings"

func intToRoman(num int) string {
	val := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := &strings.Builder{}

	for i := 0; i < 13; i++ {

		if num == 0 {
			break
		}
		times := num / val[i]

		for times > 0 {
			res.WriteString(sym[i])
			times--
		}

		num = num % val[i]

	}

	return res.String()
}
