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
