package recusion

import (
	"slices"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Two elements - sorted",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Two elements - reverse sorted",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-5, -2, 0, 3, -1},
			expected: []int{-5, -2, -1, 0, 3},
		},
		{
			name:     "Array with duplicates",
			input:    []int{5, 2, 5, 2, 1},
			expected: []int{1, 2, 2, 5, 5},
		},
		{
			name:     "Large array",
			input:    []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50},
			expected: []int{11, 12, 22, 25, 34, 45, 50, 64, 88, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the test case
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// Sort the array
			if len(input) > 0 {
				MergeSort(input, 0, len(input)-1)
			}

			// Check if the result matches expected output
			if !slices.Equal(input, tt.expected) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, input, tt.expected)
			}
		})
	}
}
