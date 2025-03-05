package main

import "math/rand"

// тут надо очень аккуратно писать удаление элемента.
// Рандомайзер будет свапать удаляемый и последний, поэтому в качестве граничных случаев:
// 1. удаляемый и есть последний, 2. удалили единственный элемент -- в обоих этих
// случаях надо просто отрезать последний элемент слайса и выйти из функции
// https://leetcode.com/problems/insert-delete-getrandom-o1/description/
type RandomizedSet struct {
	set        map[int]int
	randomizer []int
	size       int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0), 0}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = this.size // добавление ключа в мапу
	this.size++
	this.randomizer = append(this.randomizer, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}

	delete(this.set, val) // удаление ключа из мапы
	this.size--

	// метод удаления надо писать очень внимательно.
	// Крайний случай: удаляемый и был последним в рандомайзере, поэтому свапать не надо
	// Крайний случай: мы удалили единственный элемент
	if len(this.set) == 0 || index == this.size {
		this.randomizer = this.randomizer[:this.size]
		return true
	}

	last := this.randomizer[this.size]
	this.randomizer[index] = last // идея в том чтобы свапнуть удаляемый и последний а потом отрезать len
	this.set[last] = index        // строка дающая баги / ошибки, если будет исполняться на граничных случаях
	this.randomizer = this.randomizer[:this.size]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.randomizer)) // rand.Intn(n) generates a random integer in the range [0, n).
	return this.randomizer[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
