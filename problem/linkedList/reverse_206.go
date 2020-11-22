package linkedList

import "leetcodeGo/problem/model"

func reverse(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}
	var prev *model.ListNode
	var next *model.ListNode
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head = prev
	return head
}
