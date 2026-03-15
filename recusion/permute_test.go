package recusion

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// Helper to sort permutations for comparison
func sortPerms(perms [][]int) {
	for _, p := range perms {
		sort.Ints(p)
	}
	sort.Slice(perms, func(i, j int) bool {
		for k := 0; k < len(perms[i]); k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return false
	})
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "single element",
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "two elements",
			input: []int{1, 2},
			expected: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name:  "three elements",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name:  "with negative numbers",
			input: []int{-1, 0, 1},
			expected: [][]int{
				{-1, 0, 1},
				{-1, 1, 0},
				{0, -1, 1},
				{0, 1, -1},
				{1, -1, 0},
				{1, 0, -1},
			},
		},
		{
			name:  "with zero",
			input: []int{0, 1},
			expected: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name:  "empty input",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.input)

			// Validate count: n! permutations
			expectedCount := factorial(len(tt.input))
			if len(got) != expectedCount {
				t.Errorf("expected %d permutations, got %d", expectedCount, len(got))
			}

			// Sort both for order-independent comparison
			sortPerms(got)
			sortPerms(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("input=%v\nexpected=%v\ngot=%v", tt.input, tt.expected, got)
			}
		})
	}
}

// TestPermuteNoDuplicates ensures no duplicate permutations are returned
func TestPermuteNoDuplicates(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {1, 2, 3, 4}}

	for _, input := range inputs {
		got := permute(input)
		seen := make(map[string]bool)

		for _, perm := range got {
			key := fmt.Sprintf("%v", perm)
			if seen[key] {
				t.Errorf("duplicate permutation %v found for input %v", perm, input)
			}
			seen[key] = true
		}
	}
}

// TestPermuteOriginalUnchanged ensures input slice is not mutated
func TestPermuteOriginalUnchanged(t *testing.T) {
	input := []int{1, 2, 3}
	original := []int{1, 2, 3}
	permute(input)

	if !reflect.DeepEqual(input, original) {
		t.Errorf("input was mutated: expected %v, got %v", original, input)
	}
}

// TestPermuteAllElementsPresent checks every permutation has all original elements
func TestPermuteAllElementsPresent(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got := permute(input)

	inputSorted := append([]int{}, input...)
	sort.Ints(inputSorted)

	for _, perm := range got {
		permCopy := append([]int{}, perm...)
		sort.Ints(permCopy)

		if !reflect.DeepEqual(permCopy, inputSorted) {
			t.Errorf("permutation %v does not contain all elements of %v", perm, input)
		}
	}
}

// TestPermuteLengths checks every permutation has correct length
func TestPermuteLengths(t *testing.T) {
	input := []int{1, 2, 3}
	got := permute(input)

	for _, perm := range got {
		if len(perm) != len(input) {
			t.Errorf("permutation %v has length %d, expected %d", perm, len(perm), len(input))
		}
	}
}

func permutationFreq(perms [][]int) map[string]int {
	freq := make(map[string]int, len(perms))
	for _, p := range perms {
		freq[fmt.Sprintf("%v", p)]++
	}
	return freq
}

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "empty input",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name:  "all unique values",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name:  "contains duplicates",
			input: []int{1, 1, 2},
			expected: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name:  "all duplicates",
			input: []int{2, 2, 2},
			expected: [][]int{
				{2, 2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(tt.input)
			gotFreq := permutationFreq(got)
			expectedFreq := permutationFreq(tt.expected)

			if !reflect.DeepEqual(gotFreq, expectedFreq) {
				t.Errorf("input=%v\nexpected=%v\ngot=%v", tt.input, tt.expected, got)
			}
		})
	}
}

func TestPermuteUniqueSwap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "empty input",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name:  "all unique values",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name:  "contains duplicates",
			input: []int{1, 1, 2},
			expected: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name:  "all duplicates",
			input: []int{2, 2, 2},
			expected: [][]int{
				{2, 2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := append([]int{}, tt.input...)
			got := permuteUniqueSwap(inputCopy)
			gotFreq := permutationFreq(got)
			expectedFreq := permutationFreq(tt.expected)

			if !reflect.DeepEqual(gotFreq, expectedFreq) {
				t.Errorf("input=%v\nexpected=%v\ngot=%v", tt.input, tt.expected, got)
			}

			if !reflect.DeepEqual(inputCopy, tt.input) {
				t.Errorf("input was mutated: expected %v, got %v", tt.input, inputCopy)
			}
		})
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
