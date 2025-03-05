package main

// https://leetcode.com/problems/missing-number/
func missingNumber(nums []int) int {
	n := len(nums)
	sum := (1 + n) * n / 2
	for _, n := range nums {
		sum -= n
	}

	return sum
}
