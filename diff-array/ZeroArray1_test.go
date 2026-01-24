package diffarray

import (
	"testing"
)

func TestZeroArray1(t *testing.T) {
	nums := []int{2, 1, 0}
	queries := [][]int{
		{0, 1},
		{0, 2},
	}

	expected := true
	result := ZeroArray1(nums, queries)

	if result != expected {
		t.Errorf("TestZeroArray1 failed: expected %v, got %v", expected, result)
	}
}

func TestZeroArray1LargeCase(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1, 2, 1, 2}
	queries := [][]int{
		{0, 7},
		{0, 7},
		{0, 7},
		{1, 6},
		{1, 6},
		{2, 5},
		{3, 4},
	}

	expected := true
	result := ZeroArray1(nums, queries)

	if result != expected {
		t.Errorf("TestZeroArray1LargeCase failed: expected %v, got %v", expected, result)
	}
}
