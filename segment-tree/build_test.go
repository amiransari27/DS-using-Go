package segmenttree

import (
    "reflect"
    "testing"
)

func TestBuildSimple(t *testing.T) {
    nums := []int{1, 2, 3, 4}
    got := build(nums)

    expected := make([]int, 4*len(nums))
    expected[0] = 10
    expected[1] = 3
    expected[2] = 7
    expected[3] = 1
    expected[4] = 2
    expected[5] = 3
    expected[6] = 4

    if !reflect.DeepEqual(got, expected) {
        t.Fatalf("unexpected segment tree:\n got: %v\nwant: %v", got, expected)
    }
}

func TestBuildSingle(t *testing.T) {
    nums := []int{5}
    got := build(nums)

    expected := make([]int, 4*len(nums))
    expected[0] = 5

    if !reflect.DeepEqual(got, expected) {
        t.Fatalf("unexpected segment tree for single element:\n got: %v\nwant: %v", got, expected)
    }
}

func TestBuildEmptyPanics(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Fatalf("expected build to panic on empty input, but it did not")
        }
    }()

    // build should panic for empty input because it assumes non-empty slice
    _ = build([]int{})
}
