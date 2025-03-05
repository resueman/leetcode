package main

import "reflect"

// Given two strings s and t, determine if they are both one edit distance apart.
// https://wentao-shao.gitbook.io/leetcode/string/161.one-edit-distance
func IsOneEditDistance(s, t string) bool {
	var isOneEditDistance func(s, t []rune) bool // правильное использование рекурсивных анонимных
	isOneEditDistance = func(s, t []rune) bool {
		// пусть s - small
		if len(s) > len(t) {
			return isOneEditDistance(t, s)
		}

		if len(t)-len(s) > 1 {
			return false
		}

		for i := 0; i < len(s)-1; i++ { // до len(s) - 1, чтобы не выйти за границы, когда берем слайс
			if s[i] != t[i] {
				if len(s) == len(t) {
					return reflect.DeepEqual(s[i+1:], t[i+1:])
				}
				return reflect.DeepEqual(s[i:], t[i+1:])
			}
		}
		return true
	}

	s_r, t_r := []rune(s), []rune(t)
	return isOneEditDistance(s_r, t_r)
}
