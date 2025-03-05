package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// как будто ограничения из задачи могут давать неправильный ответ на интах, поэтому использую int64
// но надо еще подумать
// https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode1) bool {
	var traverse func(*TreeNode1, int64, int64) bool
	traverse = func(n *TreeNode1, min, max int64) bool {
		if n == nil {
			return true
		}

		v := int64(n.Val)
		if min < v && v < max {
			return traverse(n.Left, min, v) && traverse(n.Right, v, max)
		}
		return false
	}

	var min, max int64 = -(1 << 33), 1 << 33
	return traverse(root, min, max)
}
