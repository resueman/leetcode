package main

// https://leetcode.com/problems/binary-search/description/
func search(nums []int, target int) int {
	var getTargetIndex func(l, r int) int
	getTargetIndex = func(l, r int) int {
		if target < nums[l] && nums[r] < target ||
			l == r && nums[l] != target {
			return -1
		}

		m := (l + r) / 2
		if nums[m] == target { // это условие вместо l == r && nums[l] == target выше
			return m
		}

		i1 := getTargetIndex(l, m)
		i2 := getTargetIndex(m+1, r)
		return max(i1, i2)
	}

	if len(nums) == 0 {
		return -1
	}

	return getTargetIndex(0, len(nums)-1)
}
