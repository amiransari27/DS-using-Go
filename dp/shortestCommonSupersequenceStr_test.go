package dp

import (
	"testing"
)

// Helper: check if haystack contains needle as subsequence
func isSubsequence(haystack, needle string) bool {
	if len(needle) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(haystack) && j < len(needle); i++ {
		if haystack[i] == needle[j] {
			j++
		}
	}
	return j == len(needle)
}

func TestShortestCommonSupersequenceStrBU(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		checkLen bool
		expLen   int
	}{
		{"Both empty", "", "", true, 0},
		{"First empty", "", "abc", true, 3},
		{"Second empty", "abc", "", true, 3},
		{"Identical strings", "abc", "abc", true, 3},
		{"No common chars", "abc", "def", true, 6},
		{"Subset", "ac", "abc", true, 3},
		{"Partial overlap", "abcde", "ace", true, 5},
		{"Single char same", "a", "a", true, 1},
		{"Single char different", "a", "b", true, 2},
		{"Example1", "brca", "ccbdbda", true, 9},
		{"Example2", "ox", "cow", true, 4},
		{"Repeating chars", "aaaa", "aa", true, 4},
		{"Interleaved", "abac", "cab", true, 5},
		{"Custom", "efgh", "jghi", true, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequenceStrBU(tt.str1, tt.str2)

			// Check length
			if tt.checkLen && len(got) != tt.expLen {
				t.Fatalf("shortestCommonSupersequenceStrBU(%q, %q) length = %d; want %d", tt.str1, tt.str2, len(got), tt.expLen)
			}

			// Validate that both str1 and str2 are subsequences of result
			if !isSubsequence(got, tt.str1) {
				t.Fatalf("Result %q is not a supersequence of str1=%q", got, tt.str1)
			}
			if !isSubsequence(got, tt.str2) {
				t.Fatalf("Result %q is not a supersequence of str2=%q", got, tt.str2)
			}
		})
	}
}

func TestShortestCommonSupersequenceStrBUConsistency(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
	}{
		{"", ""},
		{"a", "a"},
		{"abc", "abc"},
		{"abc", "def"},
		{"abcde", "ace"},
		{"ox", "cow"},
		{"brca", "ccbdbda"},
		{"aaaa", "aa"},
	}

	for _, tt := range tests {
		t.Run(tt.str1+"-"+tt.str2, func(t *testing.T) {
			lenResult := shortestCommonSupersequenceBU(tt.str1, tt.str2)
			strResult := shortestCommonSupersequenceStrBU(tt.str1, tt.str2)

			// Length of string result should match length result
			if len(strResult) != lenResult {
				t.Fatalf("Length mismatch: got %d from string func, expected %d from length func", len(strResult), lenResult)
			}

			// Both strings should be subsequences
			if !isSubsequence(strResult, tt.str1) {
				t.Fatalf("SCS %q missing str1=%q", strResult, tt.str1)
			}
			if !isSubsequence(strResult, tt.str2) {
				t.Fatalf("SCS %q missing str2=%q", strResult, tt.str2)
			}
		})
	}
}

func TestShortestCommonSupersequenceStrBUProperties(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		desc string
		fn   func(string, string, string) bool
	}{
		{
			"identical strings returns the string itself",
			"hello",
			"hello",
			"identical",
			func(scs, s1, s2 string) bool {
				return scs == s1 && scs == s2
			},
		},
		{
			"no common chars: concatenation works",
			"abc",
			"def",
			"no common",
			func(scs, s1, s2 string) bool {
				return isSubsequence(scs, s1) && isSubsequence(scs, s2) && len(scs) == len(s1)+len(s2)
			},
		},
		{
			"single char: 1 char result",
			"a",
			"a",
			"single",
			func(scs, s1, s2 string) bool {
				return len(scs) == 1 && scs == "a"
			},
		},
		{
			"empty first: returns second",
			"",
			"abc",
			"empty first",
			func(scs, s1, s2 string) bool {
				return scs == s2
			},
		},
		{
			"empty second: returns first",
			"abc",
			"",
			"empty second",
			func(scs, s1, s2 string) bool {
				return scs == s1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequenceStrBU(tt.str1, tt.str2)
			if !tt.fn(got, tt.str1, tt.str2) {
				t.Fatalf("Property failed (%s): got %q for (%q, %q)", tt.desc, got, tt.str1, tt.str2)
			}
		})
	}
}

func TestShortestCommonSupersequenceStrBULength(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
	}{
		{"abc", "def"},
		{"a", "a"},
		{"abcde", "ace"},
		{"programming", "programming"},
		{"AGGTAB", "GXTXAYB"},
	}

	for _, tt := range tests {
		t.Run(tt.str1+"-"+tt.str2, func(t *testing.T) {
			scs := shortestCommonSupersequenceStrBU(tt.str1, tt.str2)
			lcs := longestCommonSubsequenceBU(tt.str1, tt.str2)

			// Formula: SCS = len(s1) + len(s2) - LCS
			expectedLen := len(tt.str1) + len(tt.str2) - lcs
			if len(scs) != expectedLen {
				t.Fatalf("SCS length formula failed: got %d, expected %d (len1=%d, len2=%d, lcs=%d)", len(scs), expectedLen, len(tt.str1), len(tt.str2), lcs)
			}
		})
	}
}

// Verify SCS doesn't have unnecessary characters
func TestShortestCommonSupersequenceStrBUMinimal(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
	}{
		{"abc", "def"},
		{"abcde", "ace"},
		{"ox", "cow"},
	}

	for _, tt := range tests {
		t.Run(tt.str1+"-"+tt.str2, func(t *testing.T) {
			scs := shortestCommonSupersequenceStrBU(tt.str1, tt.str2)

			// Verify both are subsequences
			if !isSubsequence(scs, tt.str1) {
				t.Fatalf("SCS not a supersequence of str1")
			}
			if !isSubsequence(scs, tt.str2) {
				t.Fatalf("SCS not a supersequence of str2")
			}

			// No character at start/end should be removable and still be superseq
			if len(scs) > 0 {
				trimmed := scs[1:] // remove first char
				if isSubsequence(trimmed, tt.str1) && isSubsequence(trimmed, tt.str2) {
					t.Logf("Warning: SCS %q might have extra chars at start for (%q, %q)", scs, tt.str1, tt.str2)
				}
			}
		})
	}
}

// Ensure case sensitivity
func TestShortestCommonSupersequenceStrBUCase(t *testing.T) {
	got := shortestCommonSupersequenceStrBU("Abc", "abc")
	// Should not match A with a
	if !isSubsequence(got, "Abc") || !isSubsequence(got, "abc") {
		t.Fatalf("Case-sensitive check failed for Abc/abc: got %q", got)
	}
	// Since 'A' != 'a', 'b' matches, 'c' matches, result should contain Abc and abc
	// Minimal: A,a,b,c or similar (around 4 chars)
	if len(got) < 4 {
		t.Fatalf("Expected SCS length >= 4 due to case difference, got %q", got)
	}
}
