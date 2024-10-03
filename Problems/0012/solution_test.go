package problem0012_test

import (
	"testing"

	problem0012 "github.com/realtemirov/leetcode/Problems/0012"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    int
		expected string
	}{
		{
			name:     "Test 1",
			cases:    3749,
			expected: "MMMDCCXLIX",
		},
		{
			name:     "Test 2",
			cases:    58,
			expected: "LVIII",
		},
		{
			name:     "Test 3",
			cases:    1994,
			expected: "MCMXCIV",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0012.IntToRoman(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
