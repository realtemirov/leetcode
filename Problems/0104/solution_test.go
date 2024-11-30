package problem0104_test

import (
	"testing"

	problem0104 "github.com/realtemirov/leetcode/Problems/0104"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []any
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []any{1, nil, 2},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0104.MaxDepth(helper.SliceToTree(tc.cases))
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
