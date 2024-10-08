package problem0125_test

import (
	"testing"

	problem0125 "github.com/realtemirov/leetcode/Problems/0125"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    string
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    "race a car",
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    " ",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0125.IsPalindrome(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
