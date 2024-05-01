package top150

import (
	"math/rand"
)

type RandomizedSet struct {
	list  []int
	index map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		list:  make([]int, 0),
		index: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exist := rs.index[val]; exist {
		return false
	}
	rs.list = append(rs.list, val)
	rs.index[val] = len(rs.list) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, exist := rs.index[val]
	if !exist {
		return false
	}

	lastIdx := len(rs.list) - 1
	lastVal := rs.list[lastIdx]
	rs.list[idx] = lastVal
	rs.index[lastVal] = idx
	delete(rs.index, val)
	rs.list = rs.list[:lastIdx]
	return true
}

func (rs *RandomizedSet) GetRandom() int {

	if len(rs.list) == 0 {
		return -1
	}
	return rs.list[rand.Intn(len(rs.list))]

}
