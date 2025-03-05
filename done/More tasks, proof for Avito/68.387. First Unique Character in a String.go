package main

func firstUniqChar(s string) int {
	sMap := make(map[rune]int)
	for _, c := range s {
		sMap[c]++
	}
	for i, c := range s {
		if sMap[c] == 1 {
			return i
		}
	}
	return -1
}
