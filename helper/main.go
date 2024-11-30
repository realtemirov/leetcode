package helper

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return head
}

func ListToSlice(head *ListNode) []int {
	nums := []int{}

	current := head

	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}

	return nums
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(nums) {
		current := queue[0]
		queue = queue[1:]

		// Add the left child
		if index < len(nums) && nums[index] != nil {
			current.Left = &TreeNode{Val: nums[index].(int)}
			queue = append(queue, current.Left)
		}
		index++

		// Add the right child
		if index < len(nums) && nums[index] != nil {
			current.Right = &TreeNode{Val: nums[index].(int)}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}
