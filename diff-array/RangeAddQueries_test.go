package diffarray

import (
	"testing"
)

func TestRangeAddQueries(t *testing.T) {
	tests := []struct {
		n           int
		queries     [][]int
		expected    [][]int
		description string
	}{
		{
			n: 3,
			queries: [][]int{
				{0, 0, 1, 1},
				{1, 0, 2, 2},
			},
			expected: [][]int{
				{1, 1, 0},
				{2, 2, 1},
				{1, 1, 1},
			},
			description: "3x3 matrix with 2 queries",
		},
		{
			n: 2,
			queries: [][]int{
				{0, 0, 1, 1},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
			description: "2x2 matrix with 1 query covering all",
		},
		{
			n: 4,
			queries: [][]int{
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			description: "4x4 matrix with single cell query",
		},
	}

	for _, tc := range tests {
		result := RangeAddQueries(tc.n, tc.queries)
		if !matrixEqual(result, tc.expected) {
			t.Errorf("%s: RangeAddQueries(%d, %v) = %v, expected %v",
				tc.description, tc.n, tc.queries, result, tc.expected)
		}
	}
}

func matrixEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
