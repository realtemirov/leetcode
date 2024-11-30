package problem0104

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxDepth(root *helper.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return MaxDepth(root.Right) + 1
	}

	if root.Right == nil {
		return MaxDepth(root.Left) + 1
	}

	return max(MaxDepth(root.Right), MaxDepth(root.Left)) + 1
}

func max(a, b int) int {
	if a >b {
		return a
	}

	return b
}