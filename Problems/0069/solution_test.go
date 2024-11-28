package problem0069_test

import (
	"testing"

	problem0069 "github.com/realtemirov/leetcode/Problems/0069"
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
			cases:    4,
			expected: 2,
		},
		{
			name:     "Test 2",
			cases:    8,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0069.MySqrt(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
