//1420. Build Array Where You Can Find The Maximum Exactly K Comparisons

package dp

var MOD int = 1e9 + 7

func numOfArrays(n int, m int, k int) int {
	memo := make(map[[3]int]int)
	return solveNumOfArrays(n, m, k, -1, 0, 0, memo)
}

func solveNumOfArrays(n int, m int, k int, maxVal int, maxIdx int, cost int, memo map[[3]int]int) int {

	if maxIdx == n {
		if cost == k {
			return 1
		}
		return 0
	}

	// Check memoization
	key := [3]int{maxVal, maxIdx, cost}
	if val, exists := memo[key]; exists {
		return val
	}

	result := 0
	for i := 1; i <= m; i++ {
		if maxVal < i {

			result = (result + solveNumOfArrays(n, m, k, i, maxIdx+1, cost+1, memo)) % MOD
		} else {
			result = (result + solveNumOfArrays(n, m, k, maxVal, maxIdx+1, cost, memo)) % MOD
		}
	}

	memo[key] = result % MOD
	return result % MOD
}
