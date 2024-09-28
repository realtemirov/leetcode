package problem0055_test

import (
	"testing"

	problem0055 "github.com/realtemirov/leetcode/Problems/0055"
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
			cases:    []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0055.CanJump(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
