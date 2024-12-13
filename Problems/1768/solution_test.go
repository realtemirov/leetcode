package problem1768_test

import (
	"testing"

	problem1768 "github.com/realtemirov/leetcode/Problems/1768"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected string
	}{
		{
			name:     "Test 1",
			cases:    []string{"abc", "pqr"},
			expected: "apbqcr",
		},
		{
			name:     "Test 2",
			cases:    []string{"ab", "pqrs"},
			expected: "apbqrs",
		},
		{
			name:     "Test 3",
			cases:    []string{"abcd", "pq"},
			expected: "apbqcd",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1768.MergeAlternately(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
