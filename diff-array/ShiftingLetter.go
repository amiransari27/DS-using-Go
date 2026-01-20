package diffarray

func ShiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n)

	result := []byte(s)
	// find diff array
	for _, vals := range shifts {
		l, r, dir, x := vals[0], vals[1], vals[2], 1

		if dir == 0 {
			x = -1
		}

		diff[l] += x
		if r+1 < n {
			diff[r+1] -= x
		}

	}
	//cumulative sum

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	// {(s[i] - 'a') + diff[i]} + 'a'
	// corner case |0|78|-88|
	// we have only 26 charactor
	// do wrap arond like 78%26
	// handle negative by adding 26 -> (-88%26 +26)

	// apply shift

	for i := 0; i < n; i++ {
		shift := diff[i] % 26

		if shift < 0 {
			shift += 26
		}
		result[i] = ((result[i]-'a')+byte(shift))%26 + 'a'
	}

	return string(result)
}
