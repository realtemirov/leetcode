package problem0043_test

import (
	"testing"

	problem0043 "github.com/realtemirov/leetcode/Problems/0043"
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
			cases:    []string{"2", "3"},
			expected: "6",
		},
		{
			name:     "Test 2",
			cases:    []string{"123", "456"},
			expected: "56088",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0043.Multiply(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
