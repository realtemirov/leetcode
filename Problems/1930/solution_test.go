package problem1930_test

import (
	"testing"

	problem1930 "github.com/realtemirov/leetcode/Problems/1930"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    string
		expected int
	}{
		{
			name:     "Test 1",
			cases:    "aabca",
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    "bbcbaba",
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1930.CountPalindromicSubsequence(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
