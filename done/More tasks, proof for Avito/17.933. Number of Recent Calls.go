package main

// https://leetcode.com/problems/number-of-recent-calls/description/
type RecentCounter struct {
	timestamps []int
}

func Constructor2() RecentCounter {
	return RecentCounter{make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.timestamps = append(this.timestamps, t)
	cnt := 0
	for i := len(this.timestamps) - 1; i >= 0; i-- {
		if t-this.timestamps[i] <= 3000 {
			cnt++
			continue
		}
		break
	}
	return cnt
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
