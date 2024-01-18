package leetcode

import "strings"

func IntToRoman(num int) string {
	str := strings.Builder{}

	processStr(&str, "", "M", "", num/1000)
	num = num % 1000

	processStr(&str, "M", "C", "D", num/100)
	num = num % 100

	processStr(&str, "C", "X", "L", num/10)
	num = num % 10

	processStr(&str, "X", "I", "V", num)

	return str.String()
}

	func processStr(str *strings.Builder, major string, minor string, middle string, num int) {
		if num <= 3 {
			for i := 1; i <= num; i++ {
				str.WriteString(minor)
			}
		} else if num == 4 {
			str.WriteString(minor)
			str.WriteString(middle)
		} else if num == 5 {
			str.WriteString(middle)
		} else if num <= 8 {
			str.WriteString(middle)
			for i := 6; i <= num; i++ {
				str.WriteString(minor)
			}
		} else if num == 9 {
			str.WriteString(minor)
			str.WriteString(major)
		}
	}
