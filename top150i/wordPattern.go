package top150i

import "strings"

func wordPattern(pattern string, s string) bool {

	strArr := strings.Split(s, " ")

	if len(strArr) != len(pattern) {
		return false
	}

	chToSrtMap := make(map[byte]string)
	isUsed := make(map[string]bool)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]

		if _, exist := chToSrtMap[char]; exist {
			if chToSrtMap[char] != strArr[i] {
				return false
			}

		} else {
			if isUsed[strArr[i]] {
				return false
			}
			chToSrtMap[char] = strArr[i]
			isUsed[strArr[i]] = true
		}
	}

	return true
}
