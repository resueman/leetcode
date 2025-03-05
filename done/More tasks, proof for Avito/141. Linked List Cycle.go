package main

// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && slow != fast {
		slow, fast = slow.Next, fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}

	if fast == nil {
		return false
	}
	return true
}
