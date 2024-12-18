package main

// https://leetcode.com/problems/move-zeroes/description/

// 283. Move Zeroes
// Переместить все нули в конец среза, сохранив порядок ненулевых элементов.

// Идея: заводим указатель, обозначающий место вставки ненулевого.
// 1. В цикле записываем каждый ненулевой по месту вставки.
// 2. В следующем цикле записываем нули с места вставки до конца данного слайса.
func moveZeroes(nums []int) {
	l := 0
	for _, num := range nums {
		if num != 0 {
			nums[l] = num
			l++
		}
	}

	for i := l; i < len(nums); i++ {
		nums[i] = 0
	}
}
