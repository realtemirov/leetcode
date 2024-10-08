package problem0015_test

import (
	"testing"

	problem0015 "github.com/realtemirov/leetcode/Problems/0015"
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
			cases:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Test 2",
			cases:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Test 3",
			cases:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0015.ThreeSum(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
