package main

// https://leetcode.com/problems/find-all-anagrams-in-a-string
func findAnagrams(s string, p string) []int {
	indices := make([]int, 0)
	sR, pR := []rune(s), []rune(p)
	n, k := len(sR), len(pR)
	if n < k {
		return indices
	}

	sMap, pMap := make(map[rune]int), make(map[rune]int)
	for i := 0; i < k; i++ {
		pMap[pR[i]]++
		sMap[sR[i]]++
	}

	tMatch := 0
	for k, v := range pMap {
		if sMap[k] == v {
			tMatch++ // !!!!!!!!!!!!!!
		}
	}

	expectedMatch := len(pMap) // !!!!!!!!!!!!!!
	if tMatch == expectedMatch {
		indices = append(indices, 0)
	}

	for i := k; i < n; i++ {
		in, out := sR[i], sR[i-k]
		if pMap[in] > 0 { // !!!!!!!!!!!!!!
			sMap[in]++
			if sMap[in] == pMap[in] {
				tMatch++
			} else if sMap[in]-1 == pMap[in] {
				tMatch--
			}
		}

		if pMap[out] > 0 { // !!!!!!!!!!!!!!
			sMap[out]--
			if sMap[out] == pMap[out] {
				tMatch++
			} else if sMap[out]+1 == pMap[out] {
				tMatch--
			}
		}

		if tMatch == expectedMatch {
			indices = append(indices, i-k+1)
		}
	}

	return indices
}

func findAnagrams0(s string, p string) []int {
	indices := make([]int, 0)
	if len(p) > len(s) {
		return indices
	}

	pMap := [26]int{} // arrays are comparable in go
	sMap := [26]int{}
	for i := range p {
		pMap[p[i]-'a']++
		sMap[s[i]-'a']++
	}

	l := 0
	if pMap == sMap {
		indices = append(indices, l)
	}

	for r := len(p); r < len(s); r++ {
		sMap[s[l]-'a']--
		sMap[s[r]-'a']++
		l++
		if pMap == sMap {
			indices = append(indices, l)
		}
	}

	return indices
}
