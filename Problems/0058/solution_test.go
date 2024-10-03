package problem0058_test

import (
	"testing"

	problem0058 "github.com/realtemirov/leetcode/Problems/0058"
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
			cases:    "Hello World",
			expected: 5,
		},
		{
			name:     "Test 2",
			cases:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Test 3",
			cases:    "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0058.LengthOfLastWord(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
