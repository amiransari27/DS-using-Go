package recusion

func generateParenthesis(n int) []string {

	var result []string

	var solve func([]byte)

	open := 0
	close := 0

	solve = func(s []byte) {

		if len(s) == 2*n {
			result = append(result, string(s))
			return
		}

		if open < n {
			open++
			s = append(s, '(')
			solve(s)
			s = s[:len(s)-1]
			open--
		}

		if close < open {
			close++
			s = append(s, ')')
			solve(s)
			s = s[:len(s)-1]
			close--
		}

	}

	solve([]byte{})

	return result

}
