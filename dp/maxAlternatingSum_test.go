package dp

import "testing"

func TestMaxAlternatingSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 3}, 3},
		{"Example1", []int{1, 2, 3, 4}, 4},
		{"Example2", []int{4, 2, 5, 3}, 7},
		{"Increasing", []int{1, 2, 3, 4, 5}, 5},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 5},
		{"All1", []int{1, 1, 1, 1}, 1},
		{"LargeNumbers", []int{1000, 999, 1000}, 1001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAlternatingSum(tt.nums)
			if got != tt.expected {
				t.Fatalf("MaxAlternatingSum(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestMaxAlternatingSumBU(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 3}, 3},
		{"Example1", []int{1, 2, 3, 4}, 4},
		{"Example2", []int{4, 2, 5, 3}, 7},
		{"Increasing", []int{1, 2, 3, 4, 5}, 5},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 5},
		{"All1", []int{1, 1, 1, 1}, 1},
		{"LargeNumbers", []int{1000, 999, 1000}, 1001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAlternatingSumBU(tt.nums)
			if got != tt.expected {
				t.Fatalf("MaxAlternatingSumBU(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestMaxAlternatingSumBUConstSpace(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 3}, 3},
		{"Example1", []int{1, 2, 3, 4}, 4},
		{"Example2", []int{4, 2, 5, 3}, 7},
		{"Increasing", []int{1, 2, 3, 4, 5}, 5},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 5},
		{"All1", []int{1, 1, 1, 1}, 1},
		{"LargeNumbers", []int{1000, 999, 1000}, 1001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxAlternatingSumBUConstSpace(tt.nums)
			if got != tt.expected {
				t.Fatalf("MaxAlternatingSumBUConstSpace(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
