package dp

import (
	"sort"
	"testing"
)

// Helper function to compare two sorted slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to validate if a subset is a valid divisible subset
func isValidDivisibleSubset(subset []int) bool {
	if len(subset) <= 1 {
		return true
	}
	sort.Ints(subset)
	for i := 1; i < len(subset); i++ {
		if subset[i]%subset[i-1] != 0 {
			return false
		}
	}
	return true
}

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"Empty array", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Two elements divisible", []int{1, 2}, []int{1, 2}},
		{"Two elements not divisible", []int{2, 3}, []int{2}},
		{"Example1", []int{1, 2, 3, 4, 6}, []int{1, 2, 4}},
		{"Example2", []int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
		{"Powers of 2", []int{2, 4, 8, 16}, []int{2, 4, 8, 16}},
		{"All same", []int{5, 5, 5}, []int{5, 5, 5}},
		{"No divisibility", []int{2, 3, 5, 7}, []int{2}},
		{"Complex case1", []int{1, 2, 3, 4, 6, 12}, []int{1, 2, 4, 12}},
		{"Complex case2", []int{1, 2, 3, 4, 6, 8}, []int{1, 2, 4, 8}},
		{"Single divisor", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"Unsorted input", []int{4, 1, 2, 8}, []int{1, 2, 4, 8}},
		{"With duplicates", []int{1, 2, 2, 4}, []int{1, 2, 2, 4}},
		{"Factorials", []int{1, 2, 3, 6}, []int{1, 2, 6}},
		{"Large numbers", []int{10, 20, 50, 100}, []int{10, 20, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestDivisibleSubset(tt.nums)

			// Check if lengths match
			if len(got) != len(tt.expected) {
				t.Fatalf("largestDivisibleSubset(%v) length = %d; want %d", tt.nums, len(got), len(tt.expected))
			}

			// Check if it's a valid divisible subset
			if !isValidDivisibleSubset(got) {
				t.Fatalf("largestDivisibleSubset(%v) = %v; not a valid divisible subset", tt.nums, got)
			}

			// Check if the result matches expected (order may differ due to sorting)
			if !slicesEqual(got, tt.expected) {
				t.Fatalf("largestDivisibleSubset(%v) = %v; want %v", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestLargestDivisibleSubsetProperties(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		minLen int // Minimum expected length
	}{
		{"Every element divisible by first", []int{1, 2, 4, 8, 16}, 5},
		{"Mixed", []int{3, 6, 12, 24}, 4},
		{"Powers of 3", []int{1, 3, 9, 27}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestDivisibleSubset(tt.nums)

			if len(got) < tt.minLen {
				t.Fatalf("largestDivisibleSubset(%v) length = %d; want at least %d", tt.nums, len(got), tt.minLen)
			}

			if !isValidDivisibleSubset(got) {
				t.Fatalf("largestDivisibleSubset(%v) = %v; not a valid divisible subset", tt.nums, got)
			}
		})
	}
}
