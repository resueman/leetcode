package main

func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	num, prev := m[s[0]], m[s[0]]
	for i := 1; i < len(s); i++ {
		curr := m[s[i]]
		num += curr
		if curr > prev {
			num -= 2 * prev
		}
		prev = curr
	}
	return num
}
