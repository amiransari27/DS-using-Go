package utils

import "testing"

func TestLowerBound(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		num      int
		expected int
	}{
		{
			name:     "Empty array",
			arr:      []int{},
			num:      5,
			expected: 0,
		},
		{
			name:     "Single element found",
			arr:      []int{5},
			num:      5,
			expected: 0,
		},
		{
			name:     "Single element not found",
			arr:      []int{5},
			num:      3,
			expected: 0,
		},
		{
			name:     "Single element greater",
			arr:      []int{5},
			num:      7,
			expected: 1,
		},
		{
			name:     "Element at beginning",
			arr:      []int{1, 3, 5, 7, 9},
			num:      1,
			expected: 0,
		},
		{
			name:     "Element in middle",
			arr:      []int{1, 3, 5, 7, 9},
			num:      5,
			expected: 2,
		},
		{
			name:     "Element at end",
			arr:      []int{1, 3, 5, 7, 9},
			num:      9,
			expected: 4,
		},
		{
			name:     "Element not found between values",
			arr:      []int{1, 3, 5, 7, 9},
			num:      6,
			expected: 3,
		},
		{
			name:     "Element smaller than all",
			arr:      []int{1, 3, 5, 7, 9},
			num:      0,
			expected: 0,
		},
		{
			name:     "Element larger than all",
			arr:      []int{1, 3, 5, 7, 9},
			num:      10,
			expected: 5,
		},
		{
			name:     "Duplicates present",
			arr:      []int{1, 3, 3, 3, 5, 7, 9},
			num:      3,
			expected: 1,
		},
		{
			name:     "Negative numbers",
			arr:      []int{-5, -3, -1, 0, 1, 3, 5},
			num:      -2,
			expected: 2,
		},
		{
			name:     "All same elements",
			arr:      []int{5, 5, 5, 5},
			num:      5,
			expected: 0,
		},
		{
			name:     "Large array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			num:      6,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowerBound(tt.arr, tt.num)
			if got != tt.expected {
				t.Fatalf("LowerBound(%v, %d) = %d; want %d", tt.arr, tt.num, got, tt.expected)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element found",
			arr:      []int{5},
			target:   5,
			expected: 1,
		},
		{
			name:     "Single element not found",
			arr:      []int{5},
			target:   3,
			expected: 0,
		},
		{
			name:     "Single element greater",
			arr:      []int{5},
			target:   7,
			expected: 1,
		},
		{
			name:     "Element at beginning",
			arr:      []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 1,
		},
		{
			name:     "Element in middle",
			arr:      []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 3,
		},
		{
			name:     "Element at end",
			arr:      []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 5,
		},
		{
			name:     "Element not found between values",
			arr:      []int{1, 3, 5, 7, 9},
			target:   6,
			expected: 3,
		},
		{
			name:     "Element smaller than all",
			arr:      []int{1, 3, 5, 7, 9},
			target:   0,
			expected: 0,
		},
		{
			name:     "Element larger than all",
			arr:      []int{1, 3, 5, 7, 9},
			target:   10,
			expected: 5,
		},
		{
			name:     "Duplicates present",
			arr:      []int{1, 3, 3, 3, 5, 7, 9},
			target:   3,
			expected: 4,
		},
		{
			name:     "Negative numbers",
			arr:      []int{-5, -3, -1, 0, 1, 3, 5},
			target:   -2,
			expected: 2,
		},
		{
			name:     "All same elements",
			arr:      []int{5, 5, 5, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Large array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   6,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UpperBound(tt.arr, tt.target)
			if got != tt.expected {
				t.Fatalf("UpperBound(%v, %d) = %d; want %d", tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}
