package dp

import (
"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s        string
		p        string
		expected bool
	}{
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "*",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
	}

	for _, tc := range tests {
		result := IsMatch(tc.s, tc.p)
		if result != tc.expected {
			t.Errorf("IsMatch(%q, %q) = %v, expected %v", tc.s, tc.p, result, tc.expected)
		}
	}
}
