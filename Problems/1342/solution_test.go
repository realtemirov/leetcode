package problem1342_test

import (
	"testing"

	problem1342 "github.com/realtemirov/leetcode/Problems/1342"
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
			cases:    14,
			expected: 6,
		},
		{
			name:     "Test 2",
			cases:    8,
			expected: 4,
		},
		{
			name:     "Test 3",
			cases:    123,
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1342.NumberOfSteps(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
