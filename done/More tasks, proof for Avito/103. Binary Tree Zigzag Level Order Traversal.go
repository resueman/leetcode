package main

type TreeNode4 struct {
	Val   int
	Left  *TreeNode4
	Right *TreeNode4
}

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// более короткое решение
func zigzagLevelOrder(root *TreeNode4) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	queue := []*TreeNode4{root}
	zigZagLevels := make([][]int, 0, 1)
	leftRightOrder := true
	for len(queue) != 0 {
		levelSize := len(queue) // фиксируем число нод на уровне
		levelVals := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			node := queue[0]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			// загвоздка в добавлении элемента в ответ: если уровень четный (счет с 0),
			// то добавляем как в очередь, иначе как в стек
			if leftRightOrder {
				levelVals = append(levelVals, node.Val) // queue order
			} else {
				levelVals = append([]int{node.Val}, levelVals...) // stack order
			}

			queue = queue[1:]
		}
		leftRightOrder = !leftRightOrder
		zigZagLevels = append(zigZagLevels, levelVals)
	}
	return zigZagLevels
}

func zigzagLevelOrder1(root *TreeNode4) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	queue, levels, level := []*TreeNode4{root}, []int{0}, 0
	zigZagLevels := make([][]int, 1)
	for len(queue) != 0 {
		for len(levels) > 0 && levels[0] == level {
			node := queue[0]
			if node == nil {
				queue = queue[1:]
				levels = levels[1:]
				continue
			}

			queue = append(queue, node.Left, node.Right)

			nextLevel := level + 1
			levels = append(levels, nextLevel, nextLevel)

			if level > len(zigZagLevels)-1 {
				zigZagLevels = append(zigZagLevels, make([]int, 0, 1))
			}

			// загвоздка в добавлении элемента в ответ: если уровень четный (счет с 0), то добавляем как в очередь, иначе как в стек
			if level%2 == 0 { // left traversal
				zigZagLevels[level] = append(zigZagLevels[level], node.Val) // queue order
			} else { // right traversal
				zigZagLevels[level] = append([]int{node.Val}, zigZagLevels[level]...) // stack order
			}

			queue = queue[1:]
			levels = levels[1:]
		}
		level++
	}
	return zigZagLevels
}
