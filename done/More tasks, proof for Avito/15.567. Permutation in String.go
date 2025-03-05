package main

// https://leetcode.com/problems/permutation-in-string/description/

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[rune]int)
	for _, c := range s1 {
		s1Map[c]++
	}

	s2Map := make(map[rune]int)
	windowMaxLength := len([]rune(s1))
	window := windowMaxLength
	left := 0
	s2Runes := []rune(s2)
	for r, char := range s2 {
		if _, ok := s1Map[char]; !ok {
			window = windowMaxLength
			s2Map = make(map[rune]int)
			continue
		}

		if s2Map[char]+1 > s1Map[char] {
			for ; s2Runes[left] != char; left++ {
				s2Map[s2Runes[left]]--
				window++
			}
			left++
			continue
		}

		if window == windowMaxLength {
			left = r
		}

		s2Map[char]++
		window--

		if window == 0 {
			return true
		}
	}
	return false
}

// красивое решение, разобрать перед собесом и выучить
// func checkInclusion(s1 string, s2 string) bool {
//     l, count := 0, [26]int{}
//     for i := range s1 { count[s1[i]-97]++ }

//     for r := range s2 {
//         count[s2[r]-97]--
//         if count == [26]int{} { return true }

//         if r + 1 >= len(s1) {
//             count[s2[l]-97]++
//             l++
//         }
//     }
//     return false
// }
