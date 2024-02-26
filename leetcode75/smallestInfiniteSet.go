package leetcode75

import "container/heap"

type SmallestInfiniteSet struct {
	currentSmallest int
	intHeap         *IntHeap
	intMap          map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		currentSmallest: 1,
		intHeap:         h,
		intMap:          make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	h := s.intHeap
	m := s.intMap
	if len(*h) > 0 {
		res := heap.Pop(h)
		m[res.(int)] = false
		return res.(int)
	} else {
		res := s.currentSmallest
		s.currentSmallest++
		return res
	}

}

func (s *SmallestInfiniteSet) AddBack(num int) {
	h := s.intHeap
	m := s.intMap
	if !m[num] && s.currentSmallest > num {
		m[num] = true
		heap.Push(h, num)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
