package problem3105_test

import (
	"testing"

	problem3105 "github.com/realtemirov/leetcode/Problems/3105"
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
			cases:    []int{1, 4, 3, 3, 2},
			expected: 2,
		},
		{
			name:     "Test 2",
			cases:    []int{3, 3, 3, 3},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem3105.LongestMonotonicSubarray(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
