package problem1051_test

import (
	"testing"

	problem1051 "github.com/realtemirov/leetcode/Problems/1051"
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
			cases:    []int{1, 1, 4, 2, 1, 3},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Test 3",
			cases:    []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1051.HeightChecker(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
