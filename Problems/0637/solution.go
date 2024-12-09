package problem0637

import "github.com/realtemirov/leetcode/helper"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func AverageOfLevels(root *helper.TreeNode) []float64 {
	resp := []float64{}

	queue := []*helper.TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)
		level := 0

		for i := 0; i < size; i++ {
			level += queue[i].Val

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[size:]
		resp = append(resp, float64(level)/float64(size))
	}

	return resp
}
