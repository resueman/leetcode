package main

// https://leetcode.com/problems/reverse-words-in-a-string-iii/
//Given a string s, reverse the order of characters in each word
// within a sentence while still preserving whitespace and initial word order.

// в этой задаче надо будет очень много уточнять у интервьюера:
// есть ли трейлинг пробелы, может ли быть несколько пробелов подря,
// гарантируется ли наличие хотя бы одного слова, символы из аски таблицы или
// могут быть смайлы и тп?
func reverseWord(s []rune, l, r int) {
	for l < r {
		tmp := s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	prevSpaceIndex := -1
	for i, char := range runes {
		if char == ' ' {
			reverseWord(runes, prevSpaceIndex+1, i-1)
			prevSpaceIndex = i
		}
	}
	reverseWord(runes, prevSpaceIndex+1, len(runes)-1)
	return string(runes)
}
