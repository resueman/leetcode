package main

// 1493. Longest Subarray of 1's After Deleting One Element
// Самый длинный подмассив единиц после удаления одного элемента

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/

// Идея:
// Трекаем левый (l) и правый (r) ноль.
// Считаем prevCnt -- число единиц до первого нуля, т.е. до l
// Если нуля вообще нет, возвращаем n - 1, т.к. по усл. надо удалить 1 элемент.

// Далее идем правым указателем в цикле до конца среза:
// Пытаемся обновить maxLen: currLen = prevCnt + cnt
// * если он указывает на 1, то число единиц после l: cnt = r - l.
// * если он указывает на 0, то число единиц между l и r: cnt = r - l - 1.
//   Теперь l указывает на r, prevCnt = cnt

// Вопросы:
// 1. Длина массива влезет в int?
// 2. В массиве только 0 и 1, т.е. данные гарантированно корректны?
// 3. Если в массиве только единицы, то все равно надо удалить 1 элемент? (тут да)
// Или хотим узнать макс. длину подмассива единиц после удаления одного нуля? (тут нет)

// Cлучаи:
// 1. [1], [1 1 1] --> вернуть 2, т.к. по усл. надо удалить 1 элемент
// 2. [0], [0 0], [0 0 0] --> 0
// 3. [1 1 1 0] --> 3, prevCnt = 3, второго нуля нет
// 4. [1 1 1 0 1 0] --> 4, prevCnt + cnt = 3 + 1
//    [1 1 1 0 0] --> 3, prevCnt + cnt = 3 + 0

// Сложность
// Time: O(N)
// Space: O(1)

func longestSubarray(nums []int) int {
	n := len(nums)
	var l, r int    // указывают на левый и правый нули
	var prevCnt int // считает кол-во единиц до первого нуля, т.е. до l
	for l = 0; l < n; l++ {
		if nums[l] == 0 {
			break
		}
		prevCnt++
	}

	// по усл. надо удалить 1 элемент. Т.е. если nums = [1 1 1], то ответ 2, а не 3
	if l == n {
		return n - 1
	}

	maxLen := prevCnt // не 0! Пример nums = [1 1 1 0] -- prevCnt = 3, l = 3
	for r = l + 1; r < n; r++ {
		if nums[r] == 1 {
			cnt := r - l
			maxLen = max(maxLen, prevCnt+cnt)
			continue
		}

		cnt := r - l - 1
		maxLen = max(maxLen, prevCnt+cnt)
		l, prevCnt = r, cnt
	}

	return maxLen
}
