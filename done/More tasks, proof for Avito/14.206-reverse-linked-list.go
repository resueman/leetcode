package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(arr []int) *ListNode {
	var prev *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		n := &ListNode{Val: arr[i]}
		if prev == nil {
			prev = n
			continue
		}
		n.Next = prev
		prev = n
	}
	return prev
}

func reverseListIterative(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	prev, forward := head, head.Next
	for head.Next != nil {
		head.Next = forward.Next
		forward.Next = prev
		prev = forward
		forward = head.Next
	}

	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var reverseRec func(prev *ListNode) *ListNode
	reverseRec = func(prev *ListNode) *ListNode {
		if head.Next == nil {
			return prev
		}
		forward := head.Next
		head.Next = forward.Next
		forward.Next = prev
		return reverseRec(forward)
	}

	return reverseRec(head)
}
