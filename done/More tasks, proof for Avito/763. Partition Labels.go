package main

// https://leetcode.com/problems/partition-labels/
func partitionLabels(s string) []int {
	rb := make(map[rune]int) // rb = right border
	for i, letter := range s {
		rb[letter] = i
	}

	partsLength := make([]int, 0)
	runes := []rune(s)
	lPart, rPart := 0, rb[runes[0]]
	for i := 0; i < len(runes); i++ {
		if i > rPart {
			lPart, rPart = i, rb[runes[i]]
		}

		if i == rPart {
			partsLength = append(partsLength, rPart-lPart+1)
			continue
		}

		r := rb[runes[i]]
		if r > rPart {
			rPart = r
		}
	}
	return partsLength
}

// b := borders[letter] почему-то этот код не меняет границы
// b[1] = i

// In Go, when you store an array as a value in a map, modifying the array
// does not automatically update the value in the map because arrays in Go
// have value semantics. This means that when you retrieve an array from a map,
// you are getting a copy of the array, not a reference to the actual array
// stored in the map.
