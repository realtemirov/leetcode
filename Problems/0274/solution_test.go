package problem0274_test

import (
	"testing"

	problem0274 "github.com/realtemirov/leetcode/Problems/0274"
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
			cases:    []int{3, 0, 6, 1, 5},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 3, 1},
			expected: 1,
		},
		{
			name:     "Test 3",
			cases:    []int{100},
			expected: 1,
		},
		{
			name:     "Test 4",
			cases:    []int{9, 7, 6, 2, 1},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0274.HIndex(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
