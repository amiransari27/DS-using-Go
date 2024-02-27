package leetcode75

import (
	"container/heap"
	"math"
)

func TotalCost(costs []int, k int, candidates int) int64 {
	mh1 := MinHeapTC{}
	mh2 := MinHeapTC([]int{})
	heap.Init(&mh1)
	heap.Init(&mh2)
	var cost int64
	hire := 0

	i, j := 0, len(costs)-1
	for hire < k {

		for len(mh1) < candidates && i <= j {
			heap.Push(&mh1, costs[i])
			i++
		}

		for len(mh2) < candidates && j >= i {
			heap.Push(&mh2, costs[j])
			j--
		}

		var min1 int64
		var min2 int64
		if len(mh1) > 0 {
			min1 = int64(mh1[0])
		} else {
			min1 = math.MaxInt64
		}
		if len(mh2) > 0 {
			min2 = int64(mh2[0])
		} else {
			min2 = math.MaxInt64
		}

		if min1 <= min2 {
			cost += min1
			heap.Pop(&mh1)
		} else {
			cost += min2
			heap.Pop(&mh2)
		}

		hire++
	}

	return cost
}

type MinHeapTC []int

func (m MinHeapTC) Len() int {
	return len(m)
}
func (m MinHeapTC) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m MinHeapTC) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *MinHeapTC) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *MinHeapTC) Pop() any {
	n := len(*m)
	res := (*m)[n-1]
	*m = (*m)[:n-1]
	return res
}
