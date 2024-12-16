package problem0242_test

import (
	"testing"

	problem0242 "github.com/realtemirov/leetcode/Problems/0242"
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
			cases:    []string{"anagram", "nagaram"},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []string{"rat", "car"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0242.IsAnagram(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
