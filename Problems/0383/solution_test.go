package problem0383_test

import (
	"testing"

	problem0383 "github.com/realtemirov/leetcode/Problems/0383"
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
			cases:    []string{"a", "b"},
			expected: false,
		},
		{
			name:     "Test 2",
			cases:    []string{"aa", "ab"},
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    []string{"aa", "aab"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0383.CanConstruct(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
