package top150

func StrStr(haystack string, needle string) int {

	m := len(haystack)
	n := len(needle)

	for i := 0; i <= m-n; i++ {
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
