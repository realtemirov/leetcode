package problem0151_test

import (
	"testing"

	problem0151 "github.com/realtemirov/leetcode/Problems/0151"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    string
		expected string
	}{
		{
			name:     "Test 1",
			cases:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Test 2",
			cases:    "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "Test 3",
			cases:    "a good   example",
			expected: "example good a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0151.ReverseWords(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
