package main

// https://leetcode.com/problems/matrix-diagonal-sum/
func diagonalSum(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
		sum += mat[len(mat)-i-1][i]
	}

	if len(mat)%2 == 1 { // вычесть средний, если матрица нечетная
		mid := len(mat) / 2
		sum -= mat[mid][mid]
	}

	return sum
}
