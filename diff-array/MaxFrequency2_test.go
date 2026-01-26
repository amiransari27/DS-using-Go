package diffarray

import (
	"testing"
)

func TestMaxFrequency2(t *testing.T) {
	tests := []struct {
		nums        []int
		k           int
		numOps      int
		expected    int
		description string
	}{
		{
			nums:        []int{1, 2, 4},
			k:           3,
			numOps:      2,
			expected:    3,
			description: "Convert elements within range",
		},
		{
			nums:        []int{1, 1, 1, 2},
			k:           2,
			numOps:      3,
			expected:    4,
			description: "All elements converted to same value",
		},
		{
			nums:        []int{5, 5, 5},
			k:           1,
			numOps:      2,
			expected:    3,
			description: "All same values",
		},
		{
			nums:        []int{1, 2, 3, 4, 5},
			k:           2,
			numOps:      3,
			expected:    4,
			description: "Multiple operations with range constraint",
		},
	}

	for _, tc := range tests {
		result := MaxFrequency2(tc.nums, tc.k, tc.numOps)
		if result != tc.expected {
			t.Errorf("%s: MaxFrequency2(%v, %d, %d) = %d, expected %d",
				tc.description, tc.nums, tc.k, tc.numOps, result, tc.expected)
		}
	}
}
