package leetcode

import "fmt"

var db = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans := []string{""}

	for i := 0; i < len(digits); i++ {
		char := digits[i]
		str := db[char]
		tmp := make([]string, 0)
		for j := 0; j < len(str); j++ {
			for k := 0; k < len(ans); k++ {
				tmp = append(tmp, fmt.Sprintf("%s%s", ans[k], string(str[j])))
			}
		}
		ans = tmp
	}

	return ans
}

func LetterCombinationsRec(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return solution(0, digits)
}

func solution(idx int, digits string) []string {

	if idx >= len(digits) {
		return []string{""}
	}

	ans := make([]string, 0)
	res := solution(idx+1, digits)
	char := digits[idx]
	str := db[char]
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(res); j++ {
			ans = append(ans, fmt.Sprintf("%s%s", string(str[i]), res[j]))
		}
	}
	return ans
}
