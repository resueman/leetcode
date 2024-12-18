package main

// 849. Maximize Distance to Closest Person
// https://leetcode.com/problems/maximize-distance-to-closest-person/description/

// Есть стулья, свободные и нет. Найти свободное место с максимальным расстоянием до ближайшего человека.

// Вопросы:
// 1. Может ли не быть свободных мест? (тут нет)
// 2. Могут ли все места быть свободными? (тут нет)
// 3. Может ли быть передан пустой слайс? nil слайс?

// Идея:
// 1. Есть 3 случая, а значит 3 формулы расчета maxD:
// 1 0 0 0 --> 3 = len(seats) - 1 - prev
// 0 0 0 1 --> 3 = i
// 1 0 0 1 --> 1 = (i - prev) / 2
// 1 0 1 --> 1 = (i - prev) / 2

// Сложность
// Time: O(N)
// Space: O(1)

func maxDistToClosest(seats []int) int {
	maxD := 0
	prev := -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			continue
		}

		if prev == -1 {
			maxD = max(maxD, i) // случай: крайнее левое свободное
		} else {
			maxD = max(maxD, (i-prev)/2) // случай: садимся между
		}

		prev = i
	}
	maxD = max(maxD, len(seats)-prev-1) // случай: крайнее правое свободно

	return maxD
}
