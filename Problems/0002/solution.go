package problem0002

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1, l2 *helper.ListNode) *helper.ListNode {
	head := &helper.ListNode{}
	curr := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		num := carry
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		curr.Next = &helper.ListNode{Val: num % 10}
		curr = curr.Next

		carry = num / 10
	}

	return head.Next
}
