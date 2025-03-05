package main

// https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {
	single := 0
	for _, n := range nums {
		single ^= n
	}
	return single
}
