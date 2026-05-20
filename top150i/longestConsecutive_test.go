package top150i

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Two consecutive elements",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "Two non-consecutive elements",
			nums:     []int{1, 3},
			expected: 1,
		},
		{
			name:     "Basic consecutive sequence",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Unordered consecutive sequence",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4, // [1, 2, 3, 4]
		},
		{
			name:     "With duplicates",
			nums:     []int{9, 1, 4, 7, 3, 2, 8, 5, 6},
			expected: 9, // [1, 2, 3, 4, 5, 6, 7, 8, 9]
		},
		{
			name:     "With duplicates - same numbers",
			nums:     []int{1, 2, 0, 1},
			expected: 3, // [0, 1, 2]
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, 0, 1},
			expected: 3,
		},
		{
			name:     "Negative consecutive",
			nums:     []int{-3, -2, -1, 0, 1},
			expected: 5,
		},
		{
			name:     "Multiple gaps",
			nums:     []int{1, 2, 10, 20, 30},
			expected: 2, // [1, 2]
		},
		{
			name:     "All same elements",
			nums:     []int{5, 5, 5, 5},
			expected: 1,
		},
		{
			name:     "Reverse order",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 5,
		},
		{
			name:     "Large gap in middle",
			nums:     []int{1, 2, 3, 1000, 1001, 1002},
			expected: 3, // [1000, 1001, 1002] or [1, 2, 3]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.expected {
				t.Errorf("longestConsecutive(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
