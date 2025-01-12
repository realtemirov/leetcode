package problem0226

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InvertTree(root *helper.TreeNode) *helper.TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left = InvertTree(root.Left)
	}

	if root.Right != nil {
		root.Right = InvertTree(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left

	return root
}
