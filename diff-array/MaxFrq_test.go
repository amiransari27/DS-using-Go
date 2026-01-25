package diffarray

import (
	"testing"
)

func TestMaxFrequency(t *testing.T) {
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
			description: "Convert 2 to 1 and 4 to 1",
		},
		{
			nums:        []int{1, 1, 1, 2},
			k:           2,
			numOps:      3,
			expected:    4,
			description: "All converted to 1",
		},
		{
			nums:        []int{5},
			k:           1,
			numOps:      1,
			expected:    1,
			description: "Single element",
		},
		{
			nums:        []int{1, 2, 3, 4, 5},
			k:           1,
			numOps:      2,
			expected:    3,
			description: "Multiple operations",
		},
		{
			nums:        []int{1, 4, 5},
			k:           1,
			numOps:      2,
			expected:    2,
			description: "Multiple operations",
		},
	}

	for _, tc := range tests {
		result := MaxFrequency(tc.nums, tc.k, tc.numOps)
		if result != tc.expected {
			t.Errorf("%s: MaxFrequency(%v, %d, %d) = %d, expected %d",
				tc.description, tc.nums, tc.k, tc.numOps, result, tc.expected)
		}
	}
}
