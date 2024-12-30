package problem1523_test

import (
	"testing"

	problem1523 "github.com/realtemirov/leetcode/Problems/1523"
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
			cases:    []int{3, 7},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{8, 10},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1523.CountOdds(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
