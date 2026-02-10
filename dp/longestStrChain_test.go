package dp

import "testing"

func TestLongestStrChain(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected int
	}{
		{
			name:     "Empty",
			words:    []string{},
			expected: 0,
		},
		{
			name:     "Single word",
			words:    []string{"a"},
			expected: 1,
		},
		{
			name:     "Two words, no chain",
			words:    []string{"ab", "cd"},
			expected: 1,
		},
		{
			name:     "Two words, chain",
			words:    []string{"a", "ab"},
			expected: 2,
		},
		{
			name:     "Simple chain",
			words:    []string{"a", "ab", "abc"},
			expected: 3,
		},
		{
			name:     "Example1",
			words:    []string{"a", "b", "ba", "bca", "bda", "bdca"},
			expected: 4,
		},
		{
			name:     "Example2",
			words:    []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			expected: 5,
		},
		{
			name:     "No chain possible",
			words:    []string{"abc", "def", "ghi"},
			expected: 1,
		},
		{
			name:     "Multiple chains same length",
			words:    []string{"a", "ab", "abc", "x", "xy", "xyz"},
			expected: 3,
		},
		{
			name:     "Duplicate words",
			words:    []string{"a", "a", "ab", "ab"},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestStrChain(tt.words)
			if got != tt.expected {
				t.Fatalf("longestStrChain(%v) = %d; want %d", tt.words, got, tt.expected)
			}
		})
	}
}

func TestIsPredecessor(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"a -> ab", "a", "ab", true},
		{"ab -> abc", "ab", "abc", true},
		{"ba -> bca", "ba", "bca", true},
		{"ab -> ac", "ab", "ac", false},
		{"a -> a", "a", "a", false},
		{"ab -> a", "ab", "a", false},
		{"abc -> ab", "abc", "ab", false},
		{"a -> bca", "a", "bca", false},
		{"ax -> axy", "ax", "axy", true},
		{"abc -> abcd", "abc", "abcd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPredecessor(tt.a, tt.b)
			if got != tt.expected {
				t.Fatalf("isPredecessor(%q, %q) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
