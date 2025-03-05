package main

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	l := -1
	for r := 0; r < len(nums); r++ {
		if nums[r] != val {
			if l != -1 {
				nums[l] = nums[r]
				l++
			}
			continue
		}

		// nums[r] == val
		if l == -1 {
			l = r
		}
	}

	// if no elems equal 'val'
	if l == -1 {
		return len(nums)
	}

	return l
}
