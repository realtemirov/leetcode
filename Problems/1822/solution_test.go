package problem1822_test

import (
	"testing"

	problem1822 "github.com/realtemirov/leetcode/Problems/1822"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []int{-1, -2, -3, -4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 5, 0, 2, -3},
			expected: 0,
		},
		{
			name:     "Test 3",
			cases:    []int{-1, 1, -1, 1, -1},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1822.ArraySign(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
