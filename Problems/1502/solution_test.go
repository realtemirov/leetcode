package problem1502_test

import (
	"testing"

	problem1502 "github.com/realtemirov/leetcode/Problems/1502"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    []int{3, 5, 1},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 4},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1502.CanMakeArithmeticProgression(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
