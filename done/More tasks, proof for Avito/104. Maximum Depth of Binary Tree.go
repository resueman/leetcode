package main

//https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	var getSubTreeDepth func(n *TreeNode, depth int) int
	getSubTreeDepth = func(n *TreeNode, depth int) int {
		if n == nil {
			return depth
		}
		lDepth := getSubTreeDepth(n.Left, depth+1)
		rDepth := getSubTreeDepth(n.Right, depth+1)
		return max(lDepth, rDepth)
	}
	return getSubTreeDepth(root, 0)
}
