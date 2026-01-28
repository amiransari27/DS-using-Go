package recusion

import (
	"testing"
)

func TestReverseStack(t *testing.T) {
	tests := []struct {
		input       []int
		expected    []int
		description string
	}{
		{
			input:       []int{1, 2, 3, 4, 5},
			expected:    []int{5, 4, 3, 2, 1},
			description: "Reverse stack with 5 elements",
		},
		{
			input:       []int{1},
			expected:    []int{1},
			description: "Reverse stack with 1 element",
		},
		{
			input:       []int{},
			expected:    []int{},
			description: "Reverse empty stack",
		},
		{
			input:       []int{10, 20},
			expected:    []int{20, 10},
			description: "Reverse stack with 2 elements",
		},
		{
			input:       []int{1, 2, 3},
			expected:    []int{3, 2, 1},
			description: "Reverse stack with 3 elements",
		},
	}

	for _, tc := range tests {
		result := ReverseStrack(tc.input)
		if !slicesEqual(result, tc.expected) {
			t.Errorf("%s: ReverseStrack(%v) = %v, expected %v",
				tc.description, tc.input, result, tc.expected)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
