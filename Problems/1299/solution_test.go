package problem1299_test

import (
	"testing"

	problem1299 "github.com/realtemirov/leetcode/Problems/1299"
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
			cases:    []int{17, 18, 5, 4, 6, 1},
			expected: []int{18, 6, 6, 6, 1, -1},
		},
		{
			name:     "Test 2",
			cases:    []int{400},
			expected: []int{-1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1299.ReplaceElements(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
