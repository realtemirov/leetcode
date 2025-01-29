package problem1991_test

import (
	"testing"

	problem1991 "github.com/realtemirov/leetcode/Problems/1991"
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
			cases:    []int{2, 3, -1, 8, 4},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{1, -1, 4},
			expected: 2,
		},
		{
			name:     "Test 3",
			cases:    []int{2, 5},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1991.FindMiddleIndex(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
