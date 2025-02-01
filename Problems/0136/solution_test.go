package problem0136_test

import (
	"testing"

	problem0136 "github.com/realtemirov/leetcode/Problems/0136"
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
			cases:    []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "Test 2",
			cases:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "Test 3",
			cases:    []int{1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0136.SingleNumber(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
