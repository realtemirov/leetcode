package problem0073_test

import (
	"testing"

	problem0073 "github.com/realtemirov/leetcode/Problems/0073"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    [][]int
		expected [][]int
	}{
		{
			name:     "Test 1",
			cases:    [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:     "Test 2",
			cases:    [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0073.SetZeroes(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
