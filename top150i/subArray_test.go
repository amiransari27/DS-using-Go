package top150i

import "testing"

func TestIsSubArray(t *testing.T) {
	tests := []struct {
		name     string
		arrA     []int
		arrB     []int
		expected bool
	}{
		// Test cases from the problem statement
		{
			name:     "Example 1: B is consecutive sub-array of A",
			arrA:     []int{3, 2, 7, 1, 4, 6},
			arrB:     []int{7, 1, 4},
			expected: true,
		},
		{
			name:     "Example 2: B elements in different order but consecutive in A",
			arrA:     []int{3, 2, 7, 1, 4, 6},
			arrB:     []int{4, 7, 1},
			expected: true,
		},
		{
			name:     "Example 3: B elements not consecutive in A",
			arrA:     []int{3, 2, 7, 1, 4, 6},
			arrB:     []int{3, 1, 4},
			expected: false,
		},
		{
			name:     "Example 4: B has more duplicates than consecutive window in A",
			arrA:     []int{3, 2, 7, 7, 4, 6},
			arrB:     []int{7, 4, 4},
			expected: false,
		},
		{
			name:     "Example 5: B with duplicates is sub-array of A",
			arrA:     []int{2, 2, 2, 2, 2, 7},
			arrB:     []int{2, 2, 2, 2, 7},
			expected: true,
		},

		// Edge cases
		{
			name:     "Empty B should be sub-array",
			arrA:     []int{1, 2, 3},
			arrB:     []int{},
			expected: true,
		},
		{
			name:     "Empty A and non-empty B",
			arrA:     []int{},
			arrB:     []int{1},
			expected: false,
		},
		{
			name:     "Both empty",
			arrA:     []int{},
			arrB:     []int{},
			expected: true,
		},
		{
			name:     "Single element B at start of A",
			arrA:     []int{5, 2, 3},
			arrB:     []int{5},
			expected: true,
		},
		{
			name:     "Single element B at end of A",
			arrA:     []int{5, 2, 3},
			arrB:     []int{3},
			expected: true,
		},
		{
			name:     "Single element B in middle of A",
			arrA:     []int{5, 2, 3},
			arrB:     []int{2},
			expected: true,
		},
		{
			name:     "Single element B not in A",
			arrA:     []int{5, 2, 3},
			arrB:     []int{7},
			expected: false,
		},
		{
			name:     "B length equals A length and matches",
			arrA:     []int{1, 2, 3},
			arrB:     []int{3, 1, 2},
			expected: true,
		},
		{
			name:     "B length greater than A",
			arrA:     []int{1, 2},
			arrB:     []int{1, 2, 3},
			expected: false,
		},

		// Cases with duplicates
		{
			name:     "Multiple duplicates in consecutive window",
			arrA:     []int{1, 1, 1, 2, 3},
			arrB:     []int{1, 1, 1},
			expected: true,
		},
		{
			name:     "Need exact count of duplicates",
			arrA:     []int{1, 1, 2, 3},
			arrB:     []int{1, 1, 1},
			expected: false,
		},
		{
			name:     "All same elements",
			arrA:     []int{5, 5, 5, 5},
			arrB:     []int{5, 5},
			expected: true,
		},

		// More complex cases
		{
			name:     "Sub-array at the beginning",
			arrA:     []int{1, 2, 3, 4, 5},
			arrB:     []int{3, 1, 2},
			expected: true,
		},
		{
			name:     "Sub-array at the end",
			arrA:     []int{1, 2, 3, 4, 5},
			arrB:     []int{5, 4, 3},
			expected: true,
		},
		{
			name:     "Sub-array in the middle",
			arrA:     []int{1, 2, 3, 4, 5},
			arrB:     []int{2, 3, 4},
			expected: true,
		},
		{
			name:     "Not a sub-array: elements present but not consecutive",
			arrA:     []int{1, 2, 3, 4, 5},
			arrB:     []int{1, 3, 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubArray(tt.arrA, tt.arrB)
			if result != tt.expected {
				t.Errorf("isSubArray(%v, %v) = %v, want %v", tt.arrA, tt.arrB, result, tt.expected)
			}
		})
	}
}
