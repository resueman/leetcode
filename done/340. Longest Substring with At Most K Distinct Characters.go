package main

// 340. Longest Substring with At Most K Distinct Characters
// https://algo.monster/liteproblems/340

// Найти подстроку в которой максимум k различных элементов

// Вопросы:
// 1. Что если k отрицательное?
// 2. Может ли быть дано k == 0?
// 3. k поместится в int?

// Граничные случаи:
// 1. Число уникальных элементов в мапе меньше k
// 2. k < 0
// 3. k == 0
// 4. s == ""

// Сложность
// Time: O(N)
// Space: O(K)

// Идея: заводим мапу элемент -- его кол-во, заводим left указатель, maxLen для ответа
// Идем в цикле по right указателю до длины строки, увеличивая cnt для очередного элемента
// Внутри цикла есть while (число элементов в мапе > k). Там уменьшаем cnt для элемента. Если cnt сравнялся с 0, то удаляем элемент
// В конце итерации внешнего цикла обновляем maxLen. maxLen = max(maxLen, r - l + 1)
func longestSubstrWithAtMostKDistinctChars(s string, k int) int {
	if k < 0 {
		return -1
	}

	if k == 0 {
		return 0
	}

	charCnt := make(map[byte]int)
	l := 0
	maxLen := 0
	for r := 0; r < len(s); r++ {
		charCnt[s[r]]++
		for len(charCnt) > k {
			charCnt[s[l]]--
			if charCnt[s[l]] == 0 {
				delete(charCnt, s[l])
			}
			l++
		}
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}

// println(longestSubstrWithAtMostKDistinctChars("", 5))                // 0
// println(longestSubstrWithAtMostKDistinctChars("r", -1))              // -1
// println(longestSubstrWithAtMostKDistinctChars("defaabaacabbd", 2))   // 5
// println(longestSubstrWithAtMostKDistinctChars("defaabaacabbd", 3))   // 9
// println(longestSubstrWithAtMostKDistinctChars("defffaabaacabbd", 4)) // 12
// println(longestSubstrWithAtMostKDistinctChars("defaabaacabbddd", 4)) // 12
// println(longestSubstrWithAtMostKDistinctChars("abcde", 1))           // 1
// println(longestSubstrWithAtMostKDistinctChars("abbcde", 1))          // 2
// println(longestSubstrWithAtMostKDistinctChars("abcde", 5))           // 5
// println(longestSubstrWithAtMostKDistinctChars("abcde", 15))          // 5

// То, что я написала с первой попытки за 36 минут с думаньем
// func longestSubstrWithAtMostKDistinctChars(s string, k int) int {
// 	if k < 0 {
// 		return -1
// 	}

// 	if k == 0 || len(s) == 0 {
// 		return 0
// 	}

// 	charCnt := make(map[byte]int)
// 	l, r := 0, 0
// 	maxLen := 0
// 	for r < len(s) {
// 		for r < len(s) && len(charCnt) <= k {
// 			charCnt[s[r]]++
// 			r++
// 		}

// 		if len(charCnt) > k {
// 			maxLen = r - l - 1
// 		} else {
// 			maxLen = r - l
// 		}

// 		for l < len(s) && len(charCnt) > k {
// 			out := s[l]
// 			l++
// 			charCnt[out]--
// 			if charCnt[out] == 0 {
// 				delete(charCnt, s[l])
// 				break
// 			}
// 		}
// 	}

// 	return maxLen
// }
