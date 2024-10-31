package problem0203_test

import (
	"testing"

	problem0203 "github.com/realtemirov/leetcode/Problems/0203"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			vals []int
			val  int
		}
		expected []int
	}{
		{
			name: "Test 1",
			cases: struct {
				vals []int
				val  int
			}{
				vals: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test 2",
			cases: struct {
				vals []int
				val  int
			}{
				vals: []int{},
				val:  1,
			},
			expected: []int{},
		},
		{
			name: "Test 3",
			cases: struct {
				vals []int
				val  int
			}{
				vals: []int{7, 7, 7, 7},
				val:  7,
			},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0203.RemoveElements(sliceToList(tc.cases.vals), tc.cases.val)
			require.Equal(t, tc.expected, listToSlice(result), "expected: %v, result: %v", tc.expected, result)
		})
	}
}

func sliceToList(arr []int) *problem0203.ListNode {
	if len(arr) == 0 {
		return nil
	}
	return &problem0203.ListNode{
		Val:  arr[0],
		Next: sliceToList(arr[1:]),
	}
}

func listToSlice(head *problem0203.ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append([]int{head.Val}, listToSlice(head.Next)...)
}
