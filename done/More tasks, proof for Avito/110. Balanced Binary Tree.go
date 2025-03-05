package main

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// https://leetcode.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	var getSubTreeHeight func(*TreeNode, int) (int, bool)
	getSubTreeHeight = func(n *TreeNode, height int) (int, bool) {
		if n == nil {
			return height, true
		}
		lHeight, isLeftBalanced := getSubTreeHeight(n.Left, height+1)
		rHeight, isRightBalanced := getSubTreeHeight(n.Right, height+1)
		if !isLeftBalanced || !isRightBalanced || abs(lHeight-rHeight) > 1 {
			return -1, false
		}
		return max(lHeight, rHeight), true
	}

	_, isBalanced := getSubTreeHeight(root, 0)
	return isBalanced
}
