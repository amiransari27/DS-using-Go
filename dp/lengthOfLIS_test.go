package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 1},
		{"Increasing", []int{1, 2, 3, 4, 5}, 5},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 1},
		{"Example1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"Example2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"WithDuplicates", []int{2, 2, 2, 2, 2}, 1},
		{"Mixed", []int{3, 4, -1, 0, 6, 2, 3}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLIS(tt.nums)
			if got != tt.expected {
				t.Fatalf("LengthOfLIS(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestLengthOfLISBU(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 1},
		{"Increasing", []int{1, 2, 3, 4, 5}, 5},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 1},
		{"Example1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"Example2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"WithDuplicates", []int{2, 2, 2, 2, 2}, 1},
		{"Mixed", []int{3, 4, -1, 0, 6, 2, 3}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLISBU(tt.nums)
			if got != tt.expected {
				t.Fatalf("LengthOfLISBU(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestLengthOfLISPatientSort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{5}, 1},
		{"Increasing", []int{1, 2, 3, 4, 5}, 5},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 1},
		{"Example1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"Example2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"WithDuplicates", []int{2, 2, 2, 2, 2}, 1},
		{"Mixed", []int{3, 4, -1, 0, 6, 2, 3}, 4},
		{"LargeArray", []int{1, 3, 6, 7, 9, 4, 10, 5, 5}, 6},
		{"NegativeNumbers", []int{-10, -5, 0, 5, 10}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLISPatientSort(tt.nums)
			if got != tt.expected {
				t.Fatalf("LengthOfLISPatientSort(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
