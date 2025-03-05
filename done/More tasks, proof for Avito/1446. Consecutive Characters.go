package main

// https://leetcode.com/problems/consecutive-characters/description/
func maxPower(s string) int {
	if len(s) < 2 { // в s все символы занимают 1 байт
		return len(s)
	}
	prev := s[0]
	cnt, maxCnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			cnt++
			continue
		}
		maxCnt = max(maxCnt, cnt)
		cnt, prev = 1, s[i]
	}
	maxCnt = max(cnt, maxCnt)
	return maxCnt
}
