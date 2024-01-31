package leetcode75

import "strings"

func ReverseVowels(s string) string {

	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	strArr := strings.Split(s, "")
	var low, high int = 0, len(s)

	for low < high {
		if vowels[s[low]] && vowels[s[high]] {
			tmp := strArr[high]
			strArr[high] = strArr[low]
			strArr[low] = tmp
		} else if !vowels[s[low]] {
			low++
			continue
		} else {
			high--
			continue
		}
		low++
		high--
	}

	return strings.Join(strArr, "")
}
