package recusion

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]string // if non-nil, compare exactly
		wantZero bool       // expect zero solutions
	}{
		{"n=1", 1, [][]string{{"Q"}}, false},
		{"n=2", 2, nil, true},
		{"n=3", 3, nil, true},
		{"n=4", 4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}, false},
		{"n=5", 5, nil, false}, // we don't enumerate; just check non-zero
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.n)
			if tt.wantZero {
				if len(got) != 0 {
					t.Fatalf("n=%d: expected zero solutions, got %d", tt.n, len(got))
				}
			} else if tt.expected != nil {
				if !reflect.DeepEqual(got, tt.expected) {
					t.Fatalf("got %v, want %v", got, tt.expected)
				}
			} else {
				if len(got) == 0 {
					t.Fatalf("n=%d: expected at least one solution, got none", tt.n)
				}
			}
		})
	}
}
