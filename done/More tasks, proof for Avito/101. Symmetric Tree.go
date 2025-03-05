package main

// https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode1) bool {
	var traverse func(*TreeNode1, *TreeNode1) bool
	traverse = func(n1 *TreeNode1, n2 *TreeNode1) bool {
		if n1 == nil && n2 == nil {
			return true
		}

		if n1 == nil && n2 != nil ||
			n2 == nil && n1 != nil ||
			n1.Val != n2.Val {
			return false
		}

		sym := traverse(n1.Left, n2.Right)
		if !sym {
			return false
		}
		return traverse(n1.Right, n2.Left)
	}

	l, r := root.Left, root.Right
	return traverse(l, r)
}
