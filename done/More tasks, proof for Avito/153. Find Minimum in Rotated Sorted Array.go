package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	var findIndexOfMin func(int, int) int
	findIndexOfMin = func(l, r int) int {
		if nums[l] < nums[r] {
			return l
		}

		if l+1 >= r {
			return r
		}

		mid := (l + r) / 2
		i1 := findIndexOfMin(l, mid)
		i2 := findIndexOfMin(mid+1, r)
		if nums[i1] < nums[i2] {
			return i1
		}
		return i2
	}

	i := findIndexOfMin(0, len(nums)-1)
	return nums[i]
}
