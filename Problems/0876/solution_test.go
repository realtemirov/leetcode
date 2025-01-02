package problem0876_test

import (
	"testing"

	problem0876 "github.com/realtemirov/leetcode/Problems/0876"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected []int
	}{
		{
			name:     "Test 1",
			cases:    []int{1, 2, 3, 4, 5},
			expected: []int{3, 4, 5},
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{4, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := helper.ListToSlice(problem0876.MiddleNode(helper.SliceToList(tc.cases)))
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
