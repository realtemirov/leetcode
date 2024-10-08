package problem0392_test

import (
	"testing"

	problem0392 "github.com/realtemirov/leetcode/Problems/0392"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    []string{"abc", "ahbgdc"},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []string{"axc", "ahbgdc"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0392.IsSubsequence(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
