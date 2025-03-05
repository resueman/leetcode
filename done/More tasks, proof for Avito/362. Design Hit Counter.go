package main

type HitCounter struct {
	hits [][2]int
}

func Constructor4() *HitCounter {
	return &HitCounter{make([][2]int, 0)}
}

func (c *HitCounter) Hit(timestamp int) {
	if len(c.hits) == 0 || c.hits[len(c.hits)-1][0] < timestamp {
		c.hits = append(c.hits, [2]int{timestamp, 1})
		return
	}
	c.hits[len(c.hits)-1][1]++
}

func (c *HitCounter) GetHits(timestamp int) int {
	cnt := 0
	for i := len(c.hits) - 1; i >= 0; i-- {
		if timestamp-c.hits[i][0] < 300 {
			cnt += c.hits[i][1]
			continue
		}
		break
	}
	return cnt
}

// оно вообще правильное? :)
// todo: напиши решение с кольцевым буффером, т.е. будет размер 300
// Кажется, что просто выделем массив на 300 элементов. 
// Изначально tail = 0, записывем новый элемент в tail, tail++. Если tail == 300, то tail = 0.
// Если добавляемый элемент равен element[tail - 1], то увеличиваем cnt нового, т.е. element[tail - 1][1]++
// Ну а в GetHits если буффер >= 300, то идем с конца. Иначе идем с tail до 0, а потом с i = 300 до tail  
