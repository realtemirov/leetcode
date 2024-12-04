package problem0217_test

import (
	"testing"

	problem0217 "github.com/realtemirov/leetcode/Problems/0217"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    []int{1,2,3,1},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{1,2,3,4},
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    []int{1,1,1,3,3,4,3,2,4,2},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0217.ContainsDuplicate(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
