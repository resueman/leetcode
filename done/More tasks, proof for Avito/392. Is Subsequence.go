package main

// https://leetcode.com/problems/is-subsequence/
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	i := 0
	for j := range t {
		if s[i] == t[j] {
			i++
		}

		if i == len(s) {
			return true
		}
	}

	if i == len(s) {
		return true
	}
	return false
}
