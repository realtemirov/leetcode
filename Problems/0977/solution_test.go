package problem0977_test

import (
	"testing"

	problem0977 "github.com/realtemirov/leetcode/Problems/0977"
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
			cases:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "Test 2",
			cases:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0977.SortedSquares(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
