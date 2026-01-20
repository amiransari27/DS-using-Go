package diffarray

import (
	"testing"
)

func TestShiftingLetters(t *testing.T) {
	s := "abc"
	shifts := [][]int{
		{0, 1, 0},
		{1, 2, 1},
		{0, 2, 1},
	}

	expected := "ace"
	result := ShiftingLetters(s, shifts)

	if result != expected {
		t.Errorf("TestShiftingLetters failed: expected %s, got %s", expected, result)
	}
}
