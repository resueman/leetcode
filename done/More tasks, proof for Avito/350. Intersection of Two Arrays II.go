package main

// https://leetcode.com/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	// consider nums1 < nums2 length
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]int)
	intersection := make([]int, 0)
	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if cnt, ok := m[v]; ok && cnt > 0 {
			m[v]--
			intersection = append(intersection, v)
		}
	}

	return intersection
}
