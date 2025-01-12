package problem2236

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func CheckTree(root *helper.TreeNode) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
