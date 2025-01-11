package problem0002_test

import (
	"testing"

	problem0002 "github.com/realtemirov/leetcode/Problems/0002"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    [][]int
		expected []int
	}{
		{
			name:     "Test 1",
			cases:    [][]int{{2, 4, 3}, {5, 6, 4}},
			expected: []int{7, 0, 8},
		},
		{
			name:     "Test 2",
			cases:    [][]int{{0}, {0}},
			expected: []int{0},
		},
		{
			name:     "Test 3",
			cases:    [][]int{{9, 9, 9, 9, 9, 9, 9}, {9, 9, 9, 9}},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := helper.ListToSlice(
				problem0002.AddTwoNumbers(
					helper.SliceToList(tc.cases[0]),
					helper.SliceToList(tc.cases[1]),
				),
			)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
