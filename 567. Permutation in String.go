package main

// https://leetcode.com/problems/permutation-in-string/

// Complexity:
// Space: гарантированно O(M), где M = len(s1), т.к. не добавляем в lMap элементы, отсутствующие в sMap
// Time: O(N)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s, l := s1, s2 // short, long
	sMap, lMap := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		lMap[l[i]]++
	}

	match, expectedMatch := 0, len(sMap)
	for k, v := range sMap {
		if lMap[k] == v {
			match++
		}
	}

	if match == expectedMatch {
		return true
	}

	for i := len(s); i < len(l); i++ {
		in, out := l[i], l[i-len(s)]
		if sMap[in] > 0 {
			if sMap[in] == lMap[in]+1 {
				match++
			} else if sMap[in] == lMap[in] {
				match--
			}

			lMap[in]++
		}

		if sMap[out] > 0 {
			if sMap[out] == lMap[out]-1 {
				match++
			} else if sMap[out] == lMap[out] {
				match--
			}

			lMap[out]--
		}

		if match == expectedMatch {
			return true
		}
	}

	return false
}

// неоч, тк максимальный размер мапы lMap > len(s)
func checkInclusion1(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s, l := s1, s2 // short, long
	sMap, lMap := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		lMap[l[i]]++
	}

	match, expectedMatch := 0, len(sMap)
	for k, v := range sMap {
		if lMap[k] == v {
			match++
		}
	}

	if match == expectedMatch {
		return true
	}

	for i := len(s); i < len(l); i++ {
		in, out := l[i], l[i-len(s)]
		if in == out {
			continue
		}

		if sMap[in] > 0 {
			if sMap[in] == lMap[in]+1 {
				match++
			} else if sMap[in] == lMap[in] {
				match--
			}
		}

		if sMap[out] > 0 {
			if sMap[out] == lMap[out]-1 {
				match++
			} else if sMap[out] == lMap[out] {
				match--
			}
		}

		lMap[out]--
		lMap[in]++

		if match == expectedMatch {
			return true
		}
	}

	return false
}
