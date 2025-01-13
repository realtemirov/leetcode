package problem3223_test

import (
	"testing"

	problem3223 "github.com/realtemirov/leetcode/Problems/3223"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    string
		expected int
	}{
		{
			name:     "Test 1",
			cases:    "abaacbcbb",
			expected: 5,
		},
		{
			name:     "Test 2",
			cases:    "aa",
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem3223.MinimumLength(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
