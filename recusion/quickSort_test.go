package recusion

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty", []int{}, []int{}},
		{"Single", []int{5}, []int{5}},
		{"TwoSorted", []int{1, 2}, []int{1, 2}},
		{"TwoReverse", []int{2, 1}, []int{1, 2}},
		{"AlreadySorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"ReverseSorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"WithDuplicates", []int{3, 1, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
		{"Negatives", []int{-1, -3, 2, 0}, []int{-3, -1, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			QuickSort(arr, 0, len(arr)-1)

			if !slices.Equal(arr, tt.expected) {
				t.Fatalf("QuickSort(%v) = %v; want %v", tt.input, arr, tt.expected)
			}
		})
	}
}
