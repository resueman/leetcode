package main

import (
	"sort"
)

// https://leetcode.com/problems/group-anagrams/
// Given an array of strings strs, group the anagrams together.
// You can return the answer in any order.
func groupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		// сортировка слайса рун
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sorted := string(runes)
		if _, ok := anagramsMap[sorted]; !ok {
			anagramsMap[sorted] = []string{s}
			continue
		}
		anagramsMap[sorted] = append(anagramsMap[sorted], s)
	}

	// конвертация мапы в массив ее значений
	grouped := make([][]string, 0, len(anagramsMap))
	for _, group := range anagramsMap {
		grouped = append(grouped, group)
	}
	return grouped
}
