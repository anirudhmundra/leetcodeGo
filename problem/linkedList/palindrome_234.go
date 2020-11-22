package linkedList

import "leetcodeGo/problem/model"

func isPalindrome(head *model.ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = reverse(slow)
	fast = head
	for slow != nil && fast != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
