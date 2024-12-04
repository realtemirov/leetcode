package problem0007_test

import (
	"testing"

	problem0007 "github.com/realtemirov/leetcode/Problems/0007"
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
			cases:    123,
			expected: 321,
		},
		{
			name:     "Test 2",
			cases:    -123,
			expected: -321,
		},
		{
			name:     "Test 3",
			cases:    120,
			expected: 21,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0007.Reverse(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
