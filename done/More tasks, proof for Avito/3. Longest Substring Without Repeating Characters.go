package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// можно короче, без удалений, напиши потом
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	l, curL, maxL := 0, 0, 0
	for r, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = r
			curL++
			continue
		}

		maxL = max(maxL, curL)
		index := m[v]
		for k, val := range m {
			if val < index {
				delete(m, k)
			}
		}
		m[v] = r
		l = index + 1
		curL = r - l + 1
	}
	maxL = max(maxL, curL)
	return maxL
}
