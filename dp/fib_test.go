package dp

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n           int
		expected    int
		description string
	}{
		{
			n:           0,
			expected:    0,
			description: "Fib(0) = 0",
		},
		{
			n:           1,
			expected:    1,
			description: "Fib(1) = 1",
		},
		{
			n:           2,
			expected:    1,
			description: "Fib(2) = 1",
		},
		{
			n:           5,
			expected:    5,
			description: "Fib(5) = 5",
		},
		{
			n:           10,
			expected:    55,
			description: "Fib(10) = 55",
		},
		{
			n:           15,
			expected:    610,
			description: "Fib(15) = 610",
		},
	}

	for _, tc := range tests {
		result := Fib(tc.n)
		if result != tc.expected {
			t.Errorf("%s: Fib(%d) = %d, expected %d",
				tc.description, tc.n, result, tc.expected)
		}
	}
}
func TestFibBottomUp(t *testing.T) {
	tests := []struct {
		n           int
		expected    int
		description string
	}{
		{
			n:           0,
			expected:    0,
			description: "Fib(0) = 0 (bottom-up)",
		},
		{
			n:           1,
			expected:    1,
			description: "Fib(1) = 1 (bottom-up)",
		},
		{
			n:           2,
			expected:    1,
			description: "Fib(2) = 1 (bottom-up)",
		},
		{
			n:           5,
			expected:    5,
			description: "Fib(5) = 5 (bottom-up)",
		},
		{
			n:           10,
			expected:    55,
			description: "Fib(10) = 55 (bottom-up)",
		},
		{
			n:           15,
			expected:    610,
			description: "Fib(15) = 610 (bottom-up)",
		},
		{
			n:           20,
			expected:    6765,
			description: "Fib(20) = 6765 (bottom-up)",
		},
	}

	for _, tc := range tests {
		result := FibBottomUp(tc.n)
		if result != tc.expected {
			t.Errorf("%s: Fib(%d) = %d, expected %d",
				tc.description, tc.n, result, tc.expected)
		}
	}
}
