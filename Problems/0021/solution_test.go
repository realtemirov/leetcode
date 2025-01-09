package problem0021_test

import (
	"testing"

	problem0021 "github.com/realtemirov/leetcode/Problems/0021"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			list1 []int
			list2 []int
		}
		expected []int
	}{
		{
			name: "Test 1",
			cases: struct {
				list1 []int
				list2 []int
			}{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "Test 2",
			cases: struct {
				list1 []int
				list2 []int
			}{},
			expected: []int{},
		},
		{
			name: "Test 3",
			cases: struct {
				list1 []int
				list2 []int
			}{
				list1: []int{},
				list2: []int{0},
			},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := helper.ListToSlice(
				problem0021.MergeTwoListsIterate(
					helper.SliceToList(tc.cases.list1), helper.SliceToList(tc.cases.list2),
				),
			)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
		t.Run(tc.name, func(t *testing.T) {
			result := helper.ListToSlice(
				problem0021.MergeTwoListsRecursive(
					helper.SliceToList(tc.cases.list1), helper.SliceToList(tc.cases.list2),
				),
			)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
