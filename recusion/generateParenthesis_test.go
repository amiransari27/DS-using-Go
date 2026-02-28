package recusion

import (
	"sort"
	"testing"
)

// Helper function to validate if a string has valid balanced parentheses
func isValidParenthesis(s string) bool {
	count := 0
	for _, ch := range s {
		if ch == '(' {
			count++
		} else if ch == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// Helper function to check if result contains expected strings
func slicesContainSame(got, expected []string) bool {
	if len(got) != len(expected) {
		return false
	}
	sort.Strings(got)
	sort.Strings(expected)
	for i := range got {
		if got[i] != expected[i] {
			return false
		}
	}
	return true
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{"n=0", 0, []string{""}},
		{"n=1", 1, []string{"()"}},
		{"n=2", 2, []string{"(())", "()()"}},
		{"n=3", 3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
		{"n=4", 4, nil}, // 14 combinations - skip explicit check
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)

			// Check length
			if tt.expected != nil {
				if len(got) != len(tt.expected) {
					t.Fatalf("generateParenthesis(%d) returned %d combinations, expected %d", tt.n, len(got), len(tt.expected))
				}
				// Check if same elements (order may differ)
				if !slicesContainSame(got, tt.expected) {
					t.Fatalf("generateParenthesis(%d) = %v, want %v", tt.n, got, tt.expected)
				}
			}

			// Validate each result
			for _, s := range got {
				if !isValidParenthesis(s) {
					t.Fatalf("generateParenthesis(%d) returned invalid parentheses: %q", tt.n, s)
				}
				if len(s) != 2*tt.n {
					t.Fatalf("generateParenthesis(%d) returned string of length %d, expected %d", tt.n, len(s), 2*tt.n)
				}
			}
		})
	}
}

func TestGenerateParenthesisProperties(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectedLen int // number of valid combinations (Catalan number)
	}{
		{"n=0", 0, 1},  // 1 = Catalan(0)
		{"n=1", 1, 1},  // 1 = Catalan(1)
		{"n=2", 2, 2},  // 2 = Catalan(2)
		{"n=3", 3, 5},  // 5 = Catalan(3)
		{"n=4", 4, 14}, // 14 = Catalan(4)
		{"n=5", 5, 42}, // 42 = Catalan(5)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)

			if len(got) != tt.expectedLen {
				t.Fatalf("generateParenthesis(%d) returned %d combinations, expected %d (Catalan number)", tt.n, len(got), tt.expectedLen)
			}

			// Verify all results are valid
			for _, s := range got {
				if !isValidParenthesis(s) {
					t.Fatalf("Invalid parenthesis combination: %q", s)
				}
				if len(s) != 2*tt.n {
					t.Fatalf("String length mismatch: got %d, expected %d", len(s), 2*tt.n)
				}
			}
		})
	}
}

func TestGenerateParenthesisNoDuplicates(t *testing.T) {
	n := 4
	got := generateParenthesis(n)

	// Check for duplicates
	seen := make(map[string]bool)
	for _, s := range got {
		if seen[s] {
			t.Fatalf("generateParenthesis(%d) contains duplicate: %q", n, s)
		}
		seen[s] = true
	}
}
