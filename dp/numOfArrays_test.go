package dp

import "testing"

func TestNumOfArrays(t *testing.T) {
	tests := []struct {
		name     string
		n        int // array length
		m        int // max value
		k        int // exact comparisons needed
		expected int
	}{
		{
			name:     "n=1, m=1, k=1",
			n:        1,
			m:        1,
			k:        1,
			expected: 1,
		},
		{
			name:     "n=2, m=2, k=1",
			n:        2,
			m:        2,
			k:        1,
			expected: 3,
		},
		{
			name:     "n=1, m=3, k=1",
			n:        1,
			m:        3,
			k:        1,
			expected: 3,
		},
		{
			name:     "n=3, m=1, k=1",
			n:        3,
			m:        1,
			k:        1,
			expected: 1,
		},
		{
			name:     "Example: n=2, m=3, k=1",
			n:        2,
			m:        3,
			k:        1,
			expected: 6,
		},
		{
			name:     "Example: n=5, m=2, k=3",
			n:        5,
			m:        2,
			k:        3,
			expected: 0,
		},
		{
			name:     "n=3, m=3, k=2",
			n:        3,
			m:        3,
			k:        2,
			expected: 12,
		},
		{
			name:     "n=2, m=1, k=1",
			n:        2,
			m:        1,
			k:        1,
			expected: 1,
		},
		{
			name:     "Impossible case: n=2, m=1, k=2",
			n:        2,
			m:        1,
			k:        2,
			expected: 0,
		},
		{
			name:     "n=4, m=2, k=2",
			n:        4,
			m:        2,
			k:        2,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numOfArrays(tt.n, tt.m, tt.k)
			if got != tt.expected {
				t.Fatalf("numOfArrays(%d, %d, %d) = %d; want %d", tt.n, tt.m, tt.k, got, tt.expected)
			}
		})
	}
}
