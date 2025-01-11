package problem0445

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for singly-linked list.
 * type helper.ListNode struct {
 *     Val int
 *     Next *helper.ListNode
 * }
 */
func AddTwoNumbers(l1 *helper.ListNode, l2 *helper.ListNode) *helper.ListNode {
	nums1 := make([]int, 0, 100)
	nums2 := make([]int, 0, 100)
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}
	maxL := max(len(nums1), len(nums2))
	if len(nums1) != maxL {
		nums1 = append(make([]int, maxL-len(nums1)), nums1...)
	}

	if len(nums2) != maxL {
		nums2 = append(make([]int, maxL-len(nums2)), nums2...)
	}

	carry := 0
	var head *helper.ListNode

	for i := maxL - 1; i >= 0; i-- {
		num := nums1[i] + nums2[i] + carry
		carry = num / 10

		node := &helper.ListNode{Val: num % 10}
		node.Next = head
		head = node
	}

	if carry != 0 {
		node := &helper.ListNode{Val: carry}
		node.Next = head
		head = node
	}
	if head == nil {
		return &helper.ListNode{}
	}

	return head
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
