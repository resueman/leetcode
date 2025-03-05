package main

// https://leetcode.com/problems/sum-of-left-leaves/
func sumOfLeftLeaves(root *TreeNode) int {
	var sumLeftLeaves func(*TreeNode, bool, int) int
	sumLeftLeaves = func(n *TreeNode, isLeft bool, sum int) int {
		if n.Left == nil && n.Right == nil {
			if isLeft {
				sum += n.Val
			}
			return sum
		}
		if n.Left != nil {
			sum = sumLeftLeaves(n.Left, true, sum)
		}
		if n.Right != nil {
			sum = sumLeftLeaves(n.Right, false, sum)
		}
		return sum
	}
	return sumLeftLeaves(root, false, 0)
}
