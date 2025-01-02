package problem0876

import "github.com/realtemirov/leetcode/helper"

func MiddleNode(head *helper.ListNode) *helper.ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
