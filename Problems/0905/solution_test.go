package problem0905_test

import (
	"testing"

	problem0905 "github.com/realtemirov/leetcode/Problems/0905"
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
			cases:    []int{3, 1, 2, 4},
			expected: []int{2, 4, 3, 1},
		},
		{
			name:     "Test 2",
			cases:    []int{0},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0905.SortArrayByParity(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
