package top150i

import "math/rand"

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

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.index[val]; exist {
		return false
	}

	this.list = append(this.list, val)
	this.index[val] = len(this.list) - 1
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exist := this.index[val]
	if !exist {
		return false
	}

	lastIndex := len(this.list) - 1
	lastVal := this.list[lastIndex]
	this.list[idx] = lastVal
	this.index[lastVal] = idx
	delete(this.index, val)
	this.list = this.list[:lastIndex]
	return true
}

func (this *RandomizedSet) GetRandom() int {

	if len(this.list) == 0 {
		return -1
	}

	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
