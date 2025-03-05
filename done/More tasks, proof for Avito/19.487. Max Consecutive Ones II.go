package main

// https://leetcode.ca/all/487.html
// как будто подход в том, чтобы считать единицы до первого нуля, потом единицы после него
// когда встречаем еще ноль, то складываем размеры групп единиц и смотрим максимум ли это
// когда прошли по массиву, то так же скалдываем размер предыдущей группы + текущей + 1
func GetMaxConsecutive2(a []int) int {
	prevLength, maxL := 0, 0
	curLength := 0
	l := 0
	zeroFound := false
	for r, v := range a {
		if v == 1 {
			curLength++
			continue
		}

		if !zeroFound {
			zeroFound = true
			l = r
			prevLength = l
			curLength = 0
			continue
		}

		maxL = max(maxL, prevLength+curLength+1)
		prevLength = curLength
		curLength = 0
		l = r
	}

	maxL = max(maxL, prevLength+curLength+1)
	return maxL
}
