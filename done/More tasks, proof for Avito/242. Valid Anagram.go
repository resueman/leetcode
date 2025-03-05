package main

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	for _, c := range s {
		sMap[c]++
	}

	for _, c := range t {
		if _, ok := sMap[c]; !ok {
			return false
		}
		sMap[c]--
	}

	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}

	return true
}
