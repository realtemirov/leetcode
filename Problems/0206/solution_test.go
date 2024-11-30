package problem0206_test

import (
	"testing"

	problem0206 "github.com/realtemirov/leetcode/Problems/0206"
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
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Test 3",
			cases:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0206.ReverseList(helper.SliceToList(tc.cases))
			require.Equal(t, tc.expected, helper.ListToSlice(result), "expected: %v, result: %v",
				tc.expected, helper.ListToSlice(result))
		})
	}
}
