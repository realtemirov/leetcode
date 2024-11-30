package problem0206

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *helper.ListNode) *helper.ListNode {
	var (
		prev, curr, next *helper.ListNode
	)
	curr = head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
