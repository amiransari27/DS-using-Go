package top150i

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// egg
	// add

	map1 := make(map[byte]byte)
	map2 := make(map[byte]bool)

	for i := 0; i < len(s); i++ {

		ch1 := s[i]
		ch2 := t[i]

		if _, exist := map1[ch1]; exist {

			if map1[ch1] != ch2 {
				return false
			}

		} else {
			if _, exist := map2[ch2]; exist {
				return false
			}

			map1[ch1] = ch2
			map2[ch2] = true

		}

	}

	return true

}
