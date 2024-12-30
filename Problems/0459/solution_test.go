package problem0459_test

import (
	"testing"

	problem0459 "github.com/realtemirov/leetcode/Problems/0459"
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
			cases:    "abab",
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    "aba",
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    "abcabcabcabc",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0459.RepeatedSubstringPattern(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
