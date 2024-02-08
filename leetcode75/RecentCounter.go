package leetcode75

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {

	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {

	if len(this.queue) == 0 {
		this.queue = append(this.queue, t)
	} else {
		margin := t - 3000
		for len(this.queue) > 0 && margin > this.queue[0] {
			this.queue = this.queue[1:]
		}

		this.queue = append(this.queue, t)
	}

	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
