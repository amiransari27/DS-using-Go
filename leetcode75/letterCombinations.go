package leetcode75

import "fmt"

var charMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	var solution func(ditits string) []string

	solution = func(d string) []string {
		fmt.Println("--", d)
		if len(d) == 0 {
			return []string{""}
		}
		char := d[0]
		comb := solution(d[1:])

		res := []string{}
		for _, char := range charMap[char] {
			for _, v := range comb {
				res = append(res, string(char)+v)
			}
		}

		return res
	}

	return solution(digits)
}
