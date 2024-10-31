package problem0160_test

import (
	"testing"

	problem0160 "github.com/realtemirov/leetcode/Problems/0160"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []int{},
			expected: 0,
		},
		{
			name:     "Test 2",
			cases:    []int{},
			expected: 4,
		},
		{
			name:     "Test 3",
			cases:    []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0160.Solution(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
