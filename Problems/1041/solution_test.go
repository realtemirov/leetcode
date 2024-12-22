package problem1041_test

import (
	"testing"

	problem1041 "github.com/realtemirov/leetcode/Problems/1041"
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
			cases:    "GGLLGG",
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    "GG",
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    "GL",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1041.IsRobotBounded(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
