package dp

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 1}, 2},
		{"Example1", []int{1, 2, 3, 1}, 4},
		{"Example2", []int{2, 7, 9, 3, 1}, 12},
		{"Duplicates", []int{2, 2, 2, 2}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob(tt.nums)
			if got != tt.expected {
				t.Fatalf("Rob(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestSolveBU(t *testing.T) {
	// Note: `solveBU` expects a 1-based input (index 0 is unused), so include a leading 0.
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Single", []int{0, 5}, 5},
		{"Example1", []int{0, 1, 2, 3, 1}, 4},
		{"Example2", []int{0, 2, 7, 9, 3, 1}, 12},
		{"Two", []int{0, 2, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveBU(tt.nums)
			if got != tt.expected {
				t.Fatalf("solveBU(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
