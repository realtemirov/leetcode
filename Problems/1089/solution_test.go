package problem1089_test

import (
	"testing"

	problem1089 "github.com/realtemirov/leetcode/Problems/1089"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected []int
	}{
		{
			name:     "Test 1",
			cases:    []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			problem1089.DuplicateZeros(tc.cases)
			result := tc.cases
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
