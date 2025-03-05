package main

// мой код не работает, но у него сложность O(logN) + O(K), а в эталонном решении O(log(N-K)) + O(K)
func myFindClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	closestIndex := -1
	for l <= r {
		mid := (l + r) / 2
		if l == r {
			closestIndex = mid
			break
		}

		if abs(arr[mid]-x) <= abs(arr[mid+1]-x) {
			//if x - arr[mid] <= arr[mid + k] - x {
			r = mid
		} else {
			l = mid + 1
		}
	}

	l, r = closestIndex, closestIndex
	for r-l+1 < k {
		if r == len(arr)-1 || l > 0 && abs(arr[r+1]-x) >= abs(x-arr[l-1]) {
			l--
			continue
		}
		r++
	}

	return arr[l : r+1]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// O(log(N-K)) + O(K)
// тут идея в том, чтобы сравнивать x - (левый конец левой половины) и (правый конец правой половины) - x?
// что меньше, ближе к тому и приближаемся
// вообще не понимаю, почему это работает
func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k // загвоздка в правой границе
	for l < r {
		mid := (l + r) / 2
		lD, rD := x-arr[mid], arr[mid+k]-x // основная сложность в этой строке, нельзя использовать abs
		if lD > rD {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]
}
