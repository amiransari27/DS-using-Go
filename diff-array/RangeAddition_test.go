package diffarray

import (
"testing"
)

func TestGetModifiedArray(t *testing.T) {
	length := 5
	updates := [][]int{
		{1, 3, 2},
		{2, 4, 3},
		{0, 2, -2},
	}

	expected := []int{-2, 0, 3, 5, 3}
	result := GetModifiedArray(length, updates)

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("TestGetModifiedArray failed at index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}
