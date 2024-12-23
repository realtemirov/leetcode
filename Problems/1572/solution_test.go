package problem1572_test

import (
	"testing"

	problem1572 "github.com/realtemirov/leetcode/Problems/1572"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    [][]int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: 25,
		},
		{
			name:     "Test 2",
			cases:    [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}},
			expected: 8,
		},
		{
			name:     "Test 3",
			cases:    [][]int{{5}},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1572.DiagonalSum(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
