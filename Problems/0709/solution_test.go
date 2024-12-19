package problem0709_test

import (
	"testing"

	problem0709 "github.com/realtemirov/leetcode/Problems/0709"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    string
		expected string
	}{
		{
			name:     "Test 1",
			cases:    "Hello",
			expected: "hello",
		},
		{
			name:     "Test 2",
			cases:    "here",
			expected: "here",
		},
		{
			name:     "Test 3",
			cases:    "LOVELY",
			expected: "lovely",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0709.ToLowerCase(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
