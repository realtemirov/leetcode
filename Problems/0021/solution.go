package problem0021

import "github.com/realtemirov/leetcode/helper"

func MergeTwoListsIterate(list1, list2 *helper.ListNode) *helper.ListNode {
	head := &helper.ListNode{}
	curr := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	} else if list2 == nil {
		curr.Next = list1
	}

	return head.Next
}

func MergeTwoListsRecursive(list1, list2 *helper.ListNode) *helper.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsRecursive(list2.Next, list1)
		return list2
	}
}
