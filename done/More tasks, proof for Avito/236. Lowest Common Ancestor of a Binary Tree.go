package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
// подглядела решение, так что перед собесом напиши сама

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func lowestCommonAncestor(root, p, q *TreeNode2) *TreeNode2 {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	lcaLeft := lowestCommonAncestor(root.Left, p, q)

	lcaRight := lowestCommonAncestor(root.Right, p, q)

	if lcaLeft != nil && lcaRight != nil {
		return root
	}

	if lcaLeft != nil {
		return lcaLeft
	}
	return lcaRight
}
