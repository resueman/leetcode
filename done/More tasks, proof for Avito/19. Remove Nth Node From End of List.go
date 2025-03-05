package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, r := head, head
	// число сдвигов == k
	for i := 1; i <= n; i++ {
		r = r.Next
	}

	// удаление head (k = 2, длина списка 2)
	if r == nil {
		return head.Next
	}

	// r для проверки, что пришли к элементу за которрым идет удаляемый
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next
	return head
}
