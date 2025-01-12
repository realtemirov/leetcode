package problem1480_test

import (
	"testing"

	problem1480 "github.com/realtemirov/leetcode/Problems/1480"
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
			cases:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			name:     "Test 2",
			cases:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test 3",
			cases:    []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1480.RunningSum(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
