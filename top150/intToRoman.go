package top150

import "strings"

func IntToRoman(num int) string {

	res := &strings.Builder{}

	process := func(str *strings.Builder, major byte, middle byte, minor byte, val int) {
		if val <= 3 {
			for i := 1; i <= val; i++ {
				str.WriteByte(minor)
			}
		} else if val == 4 {
			str.WriteByte(minor)
			str.WriteByte(middle)
		} else if val == 5 {
			str.WriteByte(middle)
		} else if val <= 8 {
			str.WriteByte(middle)
			for i := 6; i <= val; i++ {
				str.WriteByte(minor)
			}
		} else if val == 9 {
			str.WriteByte(minor)
			str.WriteByte(major)
		}
	}

	process(res, ' ', ' ', 'M', num/1000)
	num = num % 1000

	process(res, 'M', 'D', 'C', num/100)
	num = num % 100

	process(res, 'C', 'L', 'X', num/10)
	num = num % 10

	process(res, 'X', 'V', 'I', num)

	return res.String()
}
