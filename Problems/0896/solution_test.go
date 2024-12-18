package problem0896_test

import (
	"testing"

	problem0896 "github.com/realtemirov/leetcode/Problems/0896"
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
			cases:    []int{1, 2, 2, 3},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{6, 5, 4, 4},
			expected: true,
		},
		{
			name:     "Test 3",
			cases:    []int{1, 3, 2},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0896.IsMonotonic(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
