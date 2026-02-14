package dp

import "testing"

func TestMaxBalancedSubsequenceSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Single positive",
			nums:     []int{3},
			expected: 3,
		},
		{
			name:     "Single negative",
			nums:     []int{-5},
			expected: -5,
		},
		{
			name:     "All negative",
			nums:     []int{-1, -2, -3},
			expected: -1,
		},
		{
			name:     "All positive",
			nums:     []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{3, 3, 5, 6},
			expected: 14,
		},
		{
			name:     "With zeros",
			nums:     []int{0, 1, 2},
			expected: 3,
		},
		{
			name:     "Example with balanced subsequence",
			nums:     []int{5, -1, 5},
			expected: 5,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{10, 5, 1},
			expected: 10,
		},
		{
			name:     "Two elements valid",
			nums:     []int{1, 5},
			expected: 6,
		},
		{
			name:     "Large values",
			nums:     []int{100, -50, 100},
			expected: 100,
		},
		{
			name:     "Alternating",
			nums:     []int{2, -3, 4, -5, 6},
			expected: 12,
		},
		{
			name:     "Empty to single",
			nums:     []int{7},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxBalancedSubsequenceSum(tt.nums)
			if got != tt.expected {
				t.Fatalf("maxBalancedSubsequenceSum(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestMaxBalancedSubsequenceSumBU(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Single positive",
			nums:     []int{3},
			expected: 3,
		},
		{
			name:     "Single negative",
			nums:     []int{-5},
			expected: -5,
		},
		{
			name:     "All negative",
			nums:     []int{-1, -2, -3},
			expected: -1,
		},
		{
			name:     "All positive",
			nums:     []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{3, 3, 5, 6},
			expected: 14,
		},
		{
			name:     "With zeros",
			nums:     []int{0, 1, 2},
			expected: 3,
		},
		{
			name:     "Example with balanced subsequence",
			nums:     []int{5, -1, 5},
			expected: 5,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{10, 5, 1},
			expected: 10,
		},
		{
			name:     "Two elements valid",
			nums:     []int{1, 5},
			expected: 6,
		},
		{
			name:     "Large values",
			nums:     []int{100, -50, 100},
			expected: 100,
		},
		{
			name:     "Alternating",
			nums:     []int{2, -3, 4, -5, 6},
			expected: 12,
		},
		{
			name:     "Empty to single",
			nums:     []int{7},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxBalancedSubsequenceSumBU(tt.nums)
			if got != tt.expected {
				t.Fatalf("maxBalancedSubsequenceSumBU(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
