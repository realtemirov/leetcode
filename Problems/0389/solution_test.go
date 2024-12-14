package problem0389_test

import (
	"testing"

	problem0389 "github.com/realtemirov/leetcode/Problems/0389"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected byte
	}{
		{
			name:     "Test 1",
			cases:    []string{"abcd", "abcde"},
			expected: byte('e'),
		},
		{
			name:     "Test 2",
			cases:    []string{"", "y"},
			expected: byte('y'),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0389.FindTheDifference(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
