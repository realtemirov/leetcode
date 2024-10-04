package problem0014_test

import (
	"testing"

	problem0014 "github.com/realtemirov/leetcode/Problems/0014"
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
			cases:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Test 2",
			cases:    []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0014.LongestCommonPrefix(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
