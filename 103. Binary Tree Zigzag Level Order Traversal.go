package main

import "ya/common"

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/

func zigzagLevelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	zigzag := make([][]int, 0, 1) // !!! сделаешь len == 1, будет лишний [] в начале
	queue := []*common.TreeNode{root}
	leftToRightOrder := true
	for len(queue) > 0 {
		lvlSize := len(queue) // фиксируем кол-во элементов уровня
		lvlVals := make([]int, lvlSize)
		for i := 0; i < lvlSize; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if leftToRightOrder {
				lvlVals[i] = node.Val
			} else {
				lvlVals[lvlSize-i-1] = node.Val
			}
		}
		queue = queue[lvlSize:]
		zigzag = append(zigzag, lvlVals)
		leftToRightOrder = !leftToRightOrder
	}
	return zigzag
}
