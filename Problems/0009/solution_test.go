package problem0009_test

import (
	"testing"

	problem0009 "github.com/realtemirov/leetcode/Problems/0009"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    int
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    121,
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    -121,
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    10,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0009.IsPalindrome(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
