package problem0485_test

import (
	"testing"

	problem0485 "github.com/realtemirov/leetcode/Problems/0485"
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
			cases:    []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 0, 1, 1, 0, 1},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0485.FindMaxConsecutiveOnes(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
