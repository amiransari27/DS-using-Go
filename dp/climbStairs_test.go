package dp

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n           int
		expected    int
		description string
	}{
		{
			n:           0,
			expected:    1,
			description: "ClimbStairs(0) = 1",
		},
		{
			n:           1,
			expected:    1,
			description: "ClimbStairs(1) = 1",
		},
		{
			n:           2,
			expected:    2,
			description: "ClimbStairs(2) = 2",
		},
		{
			n:           3,
			expected:    3,
			description: "ClimbStairs(3) = 3",
		},
		{
			n:           4,
			expected:    5,
			description: "ClimbStairs(4) = 5",
		},
		{
			n:           5,
			expected:    8,
			description: "ClimbStairs(5) = 8",
		},
		{
			n:           6,
			expected:    13,
			description: "ClimbStairs(6) = 13",
		},
		{
			n:           10,
			expected:    89,
			description: "ClimbStairs(10) = 89",
		},
	}

	for _, tc := range tests {
		result := ClimbStairs(tc.n)
		if result != tc.expected {
			t.Errorf("%s: ClimbStairs(%d) = %d, expected %d",
				tc.description, tc.n, result, tc.expected)
		}
	}
}
