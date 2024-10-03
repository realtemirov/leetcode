package problem0013_test

import (
	"testing"

	problem0013 "github.com/realtemirov/leetcode/Problems/0013"
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
			cases:    "III",
			expected: 3,
		},	
		{
			name:     "Test 2",
			cases:    "LVIII",
			expected: 58,
		},
		{
			name:     "Test 3",
			cases:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0013.RomanToInt(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
