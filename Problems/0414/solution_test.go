package problem0414_test

import (
	"testing"

	problem0414 "github.com/realtemirov/leetcode/Problems/0414"
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
			cases:    []int{3, 2, 1},
			expected: 1,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2},
			expected: 2,
		},
		{
			name:     "Test 3",
			cases:    []int{2, 2, 3, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0414.ThirdMax(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
