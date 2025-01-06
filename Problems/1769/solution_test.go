package problem1769_test

import (
	"testing"

	problem1769 "github.com/realtemirov/leetcode/Problems/1769"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    string
		expected []int
	}{
		{
			name:     "Test 1",
			cases:    "110",
			expected: []int{1, 1, 3},
		},
		{
			name:     "Test 2",
			cases:    "001011",
			expected: []int{11, 8, 5, 4, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1769.MinOperations(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
