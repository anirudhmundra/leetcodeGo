package linkedList

import "leetcodeGo/problem/model"

func middleNode(head *model.ListNode) *model.ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
