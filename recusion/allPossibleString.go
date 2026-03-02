package recusion

import "sort"

func allPossibleString(s string) []string {

	var result []string

	var solve func(ba []byte, idx int)

	solve = func(ba []byte, idx int) {
		if idx >= len(s) {

			if len(ba) > 0 {
				result = append(result, string(ba))
			}

			return
		}

		ba = append(ba, s[idx])
		solve(ba, idx+1)
		ba = ba[:len(ba)-1]
		solve(ba, idx+1)
	}

	solve([]byte{}, 0)
	sort.Strings(result)
	return result
}
