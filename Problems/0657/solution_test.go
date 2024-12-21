package problem0657_test

import (
	"testing"

	problem0657 "github.com/realtemirov/leetcode/Problems/0657"
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
			cases:    "UD",
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    "LL",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0657.JudgeCircle(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
