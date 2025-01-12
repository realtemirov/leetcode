package problem0226_test

import (
	"testing"

	problem0226 "github.com/realtemirov/leetcode/Problems/0226"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []any
		expected []any
	}{
		{
			name:     "Test 1",
			cases:    []any{4, 2, 7, 1, 3, 6, 9},
			expected: []any{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Test 2",
			cases:    []any{2, 1, 3},
			expected: []any{2, 3, 1},
		},
		{
			name:     "Test 3",
			cases:    []any{},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := helper.TreeToSlice(problem0226.InvertTree(helper.SliceToTree(tc.cases)))
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
