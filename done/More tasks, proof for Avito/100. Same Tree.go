package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var isSameTreeRec func(*TreeNode, *TreeNode) bool
	isSameTreeRec = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q != nil || p != nil && q == nil {
			return false
		}

		if p == nil && q == nil {
			return true
		}

		if p.Val != q.Val {
			return false
		}

		return isSameTreeRec(p.Left, q.Left) && isSameTreeRec(p.Right, q.Right)
	}

	return isSameTreeRec(p, q)
}
