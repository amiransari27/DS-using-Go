package top150i

import "strings"

func reverseWords(s string) string {

	s = strings.TrimSpace(s)

	words := strings.Split(s, " ")

	i := 0

	for _, w := range words {
		if w != "" {
			words[i] = w
			i++
		}
	}

	words = words[:i]

	left, right := 0, len(words)-1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")

}
