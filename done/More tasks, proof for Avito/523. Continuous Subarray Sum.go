package main

// https://leetcode.com/problems/continuous-subarray-sum/
// не смогла сама догадаться, так что выучить
func checkSubarraySum(nums []int, k int) bool {
	reminders := map[int]int{0: -1} // ! не забыть инициализировать нулевым остатком с -1-м индексом
	sum := 0
	for i, n := range nums {
		sum += n
		reminder := sum % k
		if j, ok := reminders[reminder]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			reminders[reminder] = i
		}
	}
	return false
}
