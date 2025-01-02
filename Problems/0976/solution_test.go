package problem0976_test

import (
	"testing"

	problem0976 "github.com/realtemirov/leetcode/Problems/0976"
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
			cases:    []int{2,1,2},
			expected: 5,
		},
		{
			name:     "Test 2",
			cases:    []int{1,2,1,10},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0976.LargestPerimeter(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
