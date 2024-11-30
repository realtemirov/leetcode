package problem0509_test

import (
	"testing"

	problem0509 "github.com/realtemirov/leetcode/Problems/0509"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    2,
			expected: 1,
		},
		{
			name:     "Test 2",
			cases:    3,
			expected: 2,
		},
		{
			name:     "Test 3",
			cases:    4,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0509.Fib(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
