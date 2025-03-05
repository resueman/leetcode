package main

// Definition for a binary tree node.
type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

// https://leetcode.com/problems/range-sum-of-bst/
// Given the root node of a binary search tree and two integers low and high,
// return the sum of values of all nodes with a value in the inclusive range [low, high].

func rangeSumBST(root *TreeNode1, low int, high int) int {
	if root == nil {
		return 0
	}
	var traversal func(*TreeNode1, int) int
	traversal = func(n *TreeNode1, sum int) int {
		if low <= n.Val && n.Val <= high {
			sum += n.Val
		}
		if n.Left != nil {
			sum = traversal(n.Left, sum)
		}
		if n.Right != nil {
			sum = traversal(n.Right, sum)
		}
		return sum
	}
	return traversal(root, 0)
}
