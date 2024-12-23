package problem1672_test

import (
	"testing"

	problem1672 "github.com/realtemirov/leetcode/Problems/1672"
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
			cases:    [][]int{{1, 2, 3}, {3, 2, 1}},
			expected: 6,
		},
		{
			name:     "Test 2",
			cases:    [][]int{{1, 5}, {7, 3}, {3, 5}},
			expected: 10,
		},
		{
			name:     "Test 3",
			cases:    [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			expected: 17,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1672.MaximumWealth(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
