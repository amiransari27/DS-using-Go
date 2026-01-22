package diffarray

import (
	"testing"
)

func TestZeroArray(t *testing.T) {
	nums := []int{2, 0, 2}
	queries := [][]int{
		{0, 2, 1},
		{0, 2, 1},
		{1, 1, 3},
	}

	expected := 2
	result := ZeroArray(nums, queries)

	if result != expected {
		t.Errorf("TestZeroArray failed: expected %d, got %d", expected, result)
	}
}
