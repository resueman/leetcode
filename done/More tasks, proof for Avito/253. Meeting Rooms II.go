package main

import "sort"

type Interval struct {
	Start, End int
}

// Не знаю, правильно ли тут все написано, надо запускать на примерах

// Сложность O(NlogN)
// Два массива: в одном начало конференций, в другом окончания
// Сортируем их
// Начинаем идти двумя указателями по этим массивам
// Всегда увеличиваем число комнат на 1 в конце цикла
// Если текущее начало больше или равно текущего конца, уменьшаем число комнат
// Больше или равно означает освобождение одной из комнат
// Число комнат и будет ответом
// https://myinterview.guru/leetcode-253-meeting-rooms-2-52e6f0c6cae6
func MinMeetingRoomsONlogN(intervals []*Interval) int {
	start, end := make([]int, len(intervals)), make([]int, len(intervals))
	for i, interval := range intervals {
		start[i] = interval.Start
		end[i] = interval.End
	}

	sort.Slice(start, func(i, j int) bool { return start[i] < start[j] })
	sort.Slice(end, func(i, j int) bool { return end[i] < end[j] })

	s, e := 0, 0
	occupied := 0
	for s < len(start) && e < len(end) {
		occupied++
		if start[s] >= end[e] {
			occupied--
			e++
		}
		s++
	}

	return occupied
}

// Сложность O(N), если конференции проводятся в небольшой временной промежуток
// Алгоритм: смотрим, как меняется delta, т.е. кол-во занятых / свободных комнат в момент t
// Например в момент t начались 2 конференции, тогда delta[t] = +2.
// Если закончилась одна конфа, тогда delta[t] = -1.
// Если одна конфа закончилась, одна началась, то delta[t] = 0.
// Далее считаем префикс сумму t, получаем сколько комнат занято в момент t
// Максимальное значение префикс суммы и есть число необходимых комнат
// https://algo.monster/liteproblems/253
func MinMeetingRoomsON(intervals []*Interval) int {
	maxT := intervals[0].End
	for _, i := range intervals {
		maxT = max(maxT, i.End)
	}

	delta := make([]int, maxT+1)
	for _, i := range intervals {
		delta[i.Start]++
		delta[i.End]--
	}

	occupied := 0
	for i := 1; i < len(delta); i++ {
		delta[i] += delta[i-1]
		occupied = max(occupied, delta[i])
	}

	return occupied
}
