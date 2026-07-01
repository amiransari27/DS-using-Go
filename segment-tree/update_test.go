package segmenttree

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	st := build(nums)

	update(2, 7, 0, 0, len(nums)-1, &st, nums)

	if nums[2] != 7 {
		t.Fatalf("expected nums[2] to be updated to 7, got %d", nums[2])
	}

	expected := make([]int, len(st))
	expected[0] = 14
	expected[1] = 3
	expected[2] = 11
	expected[3] = 1
	expected[4] = 2
	expected[5] = 7
	expected[6] = 4

	if !reflect.DeepEqual(st, expected) {
		t.Fatalf("unexpected segment tree after update:\n got: %v\nwant: %v", st, expected)
	}
}
