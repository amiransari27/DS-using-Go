package dp

import "testing"

func TestRob2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 1}, 2},
		{"Example1", []int{1, 2, 3, 1}, 4},
		{"Example2", []int{2, 7, 9, 3, 1}, 11},
		{"CircularCase", []int{2, 3, 2}, 3},
		{"Duplicates", []int{2, 2, 2, 2}, 4},
		{"LargeLast", []int{1, 3, 1, 3, 100}, 103},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob2(tt.nums)
			if got != tt.expected {
				t.Fatalf("Rob2(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestRob2BU(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 1}, 2},
		{"Example1", []int{1, 2, 3, 1}, 4},
		{"Example2", []int{2, 7, 9, 3, 1}, 11},
		{"CircularCase", []int{2, 3, 2}, 3},
		{"Duplicates", []int{2, 2, 2, 2}, 4},
		{"LargeLast", []int{1, 3, 1, 3, 100}, 103},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob2BU(tt.nums)
			if got != tt.expected {
				t.Fatalf("Rob2BU(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestRob2BUConstSpace(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 5},
		{"Two", []int{2, 1}, 2},
		{"Example1", []int{1, 2, 3, 1}, 4},
		{"Example2", []int{2, 7, 9, 3, 1}, 11},
		{"CircularCase", []int{2, 3, 2}, 3},
		{"Duplicates", []int{2, 2, 2, 2}, 4},
		{"LargeLast", []int{1, 3, 1, 3, 100}, 103},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob2BUConstSpace(tt.nums)
			if got != tt.expected {
				t.Fatalf("Rob2BUConstSpace(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
