package dp

import "testing"

func TestShortestCommonSupersequence(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected int
	}{
		{"Both empty", "", "", 0},
		{"First empty", "", "abc", 3},
		{"Second empty", "abc", "", 3},
		{"Identical strings", "abc", "abc", 3},
		{"No common chars", "abc", "def", 6},
		{"Subset", "ac", "abc", 3},
		{"Partial overlap", "abcde", "ace", 5},
		{"Single char same", "a", "a", 1},
		{"Single char different", "a", "b", 2},
		{"Example1", "brca", "ccbdbda", 9},
		{"Example2", "ox", "cow", 4},
		{"Repeating chars", "aaaa", "aa", 4},
		{"Interleaved", "abac", "cab", 5},
		{"Custom", "efgh", "jghi", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequence(tt.str1, tt.str2)
			if got != tt.expected {
				t.Fatalf("shortestCommonSupersequence(%q, %q) = %d; want %d", tt.str1, tt.str2, got, tt.expected)
			}
		})
	}
}

func TestShortestCommonSupersequenceBU(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected int
	}{
		{"Both empty", "", "", 0},
		{"First empty", "", "abc", 3},
		{"Second empty", "abc", "", 3},
		{"Identical strings", "abc", "abc", 3},
		{"No common chars", "abc", "def", 6},
		{"Subset", "ac", "abc", 3},
		{"Partial overlap", "abcde", "ace", 5},
		{"Single char same", "a", "a", 1},
		{"Single char different", "a", "b", 2},
		{"Example1", "brca", "ccbdbda", 9},
		{"Example2", "ox", "cow", 4},
		{"Repeating chars", "aaaa", "aa", 4},
		{"Interleaved", "abac", "cab", 5},
		{"Custom", "efgh", "jghi", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequenceBU(tt.str1, tt.str2)
			if got != tt.expected {
				t.Fatalf("shortestCommonSupersequenceBU(%q, %q) = %d; want %d", tt.str1, tt.str2, got, tt.expected)
			}
		})
	}
}

func TestShortestCommonSupersequenceProperties(t *testing.T) {
	tests := []struct {
		name  string
		str1  string
		str2  string
		check func(int, string, string) bool
	}{
		{
			"SCS length >= max(len(str1), len(str2))",
			"abc",
			"def",
			func(scs int, s1, s2 string) bool {
				return scs >= len(s1) && scs >= len(s2)
			},
		},
		{
			"SCS length <= len(str1) + len(str2)",
			"abcde",
			"ace",
			func(scs int, s1, s2 string) bool {
				return scs <= len(s1)+len(s2)
			},
		},
		{
			"Identical strings: SCS = length of string",
			"hello",
			"hello",
			func(scs int, s1, s2 string) bool {
				return scs == len(s1)
			},
		},
		{
			"No common chars: SCS = len(s1) + len(s2)",
			"abc",
			"def",
			func(scs int, s1, s2 string) bool {
				return scs == len(s1)+len(s2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequence(tt.str1, tt.str2)
			if !tt.check(got, tt.str1, tt.str2) {
				t.Fatalf("Property check failed for (%q, %q): SCS = %d", tt.str1, tt.str2, got)
			}
		})
	}
}

// Verify SCS = len(s1) + len(s2) - LCS using formula
func TestShortestCommonSupersequenceFormula(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
	}{
		{"abcde", "ace"},
		{"ox", "cow"},
		{"AGGTAB", "GXTXAYB"},
		{"abcd", "dcba"},
	}

	for _, tt := range tests {
		t.Run(tt.str1+"-"+tt.str2, func(t *testing.T) {
			scs := shortestCommonSupersequence(tt.str1, tt.str2)
			lcs := longestCommonSubsequenceBU(tt.str1, tt.str2)
			// Formula: SCS = len(s1) + len(s2) - LCS
			expected := len(tt.str1) + len(tt.str2) - lcs
			if scs != expected {
				t.Fatalf("SCS formula check failed: got %d, want %d (len1=%d, len2=%d, lcs=%d)", scs, expected, len(tt.str1), len(tt.str2), lcs)
			}
		})
	}
}
