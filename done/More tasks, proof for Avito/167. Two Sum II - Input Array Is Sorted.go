package main

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum1(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] > target {
			r--
			continue
		}
		l++
	}
	return []int{}
}
