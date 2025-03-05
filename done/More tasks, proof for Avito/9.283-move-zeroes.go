package main

// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
func moveZeroes(nums []int) {
	l := 0
	for _, n := range nums {
		if n != 0 {
			nums[l] = n
			l++
		}
	}
	for i := l; i < len(nums); i++ {
		nums[i] = 0
	}
}
