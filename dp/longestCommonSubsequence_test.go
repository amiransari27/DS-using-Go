package dp

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{"Both empty", "", "", 0},
		{"First empty", "", "abc", 0},
		{"Second empty", "abc", "", 0},
		{"Identical", "abc", "abc", 3},
		{"No common", "abc", "def", 0},
		{"Example1", "abcde", "ace", 3},
		{"Repeats", "aaaa", "aa", 2},
		{"Interleaved", "abac", "cab", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.expected {
				t.Fatalf("longestCommonSubsequence(%q, %q) = %d; want %d", tt.text1, tt.text2, got, tt.expected)
			}
		})
	}
}

func TestLongestCommonSubsequenceBU(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{"Both empty", "", "", 0},
		{"First empty", "", "abc", 0},
		{"Second empty", "abc", "", 0},
		{"Identical", "abc", "abc", 3},
		{"No common", "abc", "def", 0},
		{"Example1", "abcde", "ace", 3},
		{"Repeats", "aaaa", "aa", 2},
		{"Interleaved", "abac", "cab", 2},
		{"Single char", "a", "a", 1},
		{"Single char no match", "a", "b", 0},
		{"LeetCode example1", "ox", "cow", 1},
		{"LeetCode example2", "abc", "abc", 3},
		{"LeetCode example3", "abc", "def", 0},
		{"Longer strings", "oxcpqrsvwf", "sxyspmqyhbt", 3},
		{"All common", "programming", "programming", 11},
		{"Partial overlap", "AGGTAB", "GXTXAYB", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequenceBU(tt.text1, tt.text2)
			if got != tt.expected {
				t.Fatalf("longestCommonSubsequenceBU(%q, %q) = %d; want %d", tt.text1, tt.text2, got, tt.expected)
			}
		})
	}
}

func TestLongestCommonSubsequenceStringBU(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected string
	}{
		{"Both empty", "", "", ""},
		{"First empty", "", "abc", ""},
		{"Second empty", "abc", "", ""},
		{"Identical", "abc", "abc", "abc"},
		{"No common", "abc", "def", ""},
		{"Example1", "abcde", "ace", "ace"},
		{"Repeats", "aaaa", "aa", "aa"},
		{"Interleaved", "abac", "cab", "ab"},
		{"Single char", "a", "a", "a"},
		{"Single char no match", "a", "b", ""},
		{"LeetCode example1", "ox", "cow", "o"},
		{"Partial overlap", "AGGTAB", "GXTXAYB", "GTAB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequenceStringBU(tt.text1, tt.text2)
			if got != tt.expected {
				t.Fatalf("longestCommonSubsequenceStringBU(%q, %q) = %q; want %q", tt.text1, tt.text2, got, tt.expected)
			}
		})
	}
}
