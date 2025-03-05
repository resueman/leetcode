package main

// Given an array of integers nums and an integer k,
// return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.
// https://leetcode.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	sum, cnt := 0, 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // важный момент, до которого надо догадаться во время дебага на крайних случаях
	for _, n := range nums {
		sum += n
		cnt += prefixSum[sum-k] // важно не перепутать эту и следующую строку
		prefixSum[sum]++
	}
	return cnt
	// надо знать как мапа работает:
	// что возвращается default of value type, т.е. дефолт int равный 0
	// если нет значение и делаем ++, то все ок
}
