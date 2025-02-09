package problem2610_test

import (
	"testing"

	problem2610 "github.com/realtemirov/leetcode/Problems/2610"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected [][]int
	}{
		{
			name:     "Test 1",
			cases:    []int{1, 3, 4, 1, 2, 3, 1},
			expected: [][]int{{1, 3, 4, 2}, {1, 3}, {1}},
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 3, 4},
			expected: [][]int{{1, 2, 3, 4}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2610.FindMatrix(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
