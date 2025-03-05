package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	left, curr, uniqueCnt := 1, nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != curr {
			nums[left] = nums[i]
			curr = nums[i]
			left++
			uniqueCnt++
		}
	}
	return uniqueCnt
}
