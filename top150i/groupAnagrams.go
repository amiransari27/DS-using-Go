package top150i

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	res := make([][]string, 0)

	mp := make(map[string][]string)

	for _, str := range strs {

		chars := strings.Split(str, "")
		sort.Strings(chars)
		tmpStr := strings.Join(chars, "")

		mp[tmpStr] = append(mp[tmpStr], str)

	}

	for _, strs := range mp {

		res = append(res, strs)
	}

	return res

}

func groupAnagrams2(strs []string) [][]string {

	res := make([][]string, 0)

	mp := make(map[string][]string)

	var putToMap func(str string)

	putToMap = func(str string) {

		charArr := [26]int{}

		for i := 0; i < len(str); i++ {
			charArr[str[i]-'a']++
		}

		chars := make([]byte, 0)
		for i := 0; i < 26; i++ {
			char := i + 'a'
			for j := 0; j < charArr[i]; j++ {
				chars = append(chars, byte(char))
			}
		}

		mp[string(chars)] = append(mp[string(chars)], str)

	}

	for _, str := range strs {
		putToMap(str)
	}

	for _, strs := range mp {

		res = append(res, strs)
	}

	return res

}
