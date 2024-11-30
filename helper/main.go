package helper

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func NewUtilsListNode(val int) *ListNode {
// 	return &ListNode{Val: val}
// }

// // NextSetter allows setting the next node in the list.
// type NextSetter interface {
// 	SetNext(interface{ NextSetter })
// }

// // Implement SetNext for *utils.ListNode
// func (n *ListNode) SetNext(next interface{ NextSetter }) {
// 	n.Next = next.(*ListNode)
// }

// // SliceToList creates a linked list from a slice of integers using a factory function.
// func SliceToList(nums []int, newNode func(int) interface{ NextSetter }) interface{ NextSetter } {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	head := newNode(nums[0])
// 	current := head
// 	for _, val := range nums[1:] {
// 		newNode := newNode(val)
// 		current.SetNext(newNode)
// 		current = newNode
// 	}
// 	return head
// }

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
