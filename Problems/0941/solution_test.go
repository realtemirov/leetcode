package problem0941_test

import (
	"testing"

	problem0941 "github.com/realtemirov/leetcode/Problems/0941"
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
			cases:    []int{2, 1},
			expected: false,
		},
		{
			name:     "Test 2",
			cases:    []int{3, 5, 5},
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    []int{0, 3, 2, 1},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0941.ValidMountainArray(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
