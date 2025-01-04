package problem2270_test

import (
	"testing"

	problem2270 "github.com/realtemirov/leetcode/Problems/2270"
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
			cases:    []int{10,4,-8,7},
			expected: 2,
		},
		{
			name:     "Test 2",
			cases:    []int{2,3,1,0},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2270.WaysToSplitArray(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
