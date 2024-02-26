package leetcode75

import (
	"container/heap"
	"sort"
)

func MaxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	vector := make([][]int, n)

	for i, val := range nums1 {
		vector[i] = []int{val, nums2[i]}
	}
	sort.Slice(vector, func(i, j int) bool {
		return vector[i][1] > vector[j][1]
	})

	h := MinHeap([]int{})
	heap.Init(&h)
	var ksum int64 = 0

	for i := 0; i < k; i++ {
		ksum += int64(vector[i][0])
		heap.Push(&h, vector[i][0])
	}
	var res int64 = ksum * int64(vector[k-1][1])

	for i := k; i < n; i++ {
		ksum += int64(vector[i][0]) - int64(heap.Pop(&h).(int))

		heap.Push(&h, vector[i][0])

		res = max(res, ksum*int64(vector[i][1]))
	}

	return res
}

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}
func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() any {
	n := len(*m)
	res := (*m)[n-1]
	*m = (*m)[:n-1]
	return res
}
