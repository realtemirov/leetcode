package problem3042_test

import (
	"testing"

	problem3042 "github.com/realtemirov/leetcode/Problems/3042"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []string{"a","aba","ababa","aa"},
			expected: 4,
		},
		{
			name:     "Test 2",
			cases:    []string{"pa","papa","ma","mama"},
			expected: 2,
		},
		{
			name:     "Test 3",
			cases:    []string{"abab","ab"},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem3042.CountPrefixSuffixPairs(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
