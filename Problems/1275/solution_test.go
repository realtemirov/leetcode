package problem1275_test

import (
	"testing"

	problem1275 "github.com/realtemirov/leetcode/Problems/1275"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    [][]int
		expected string
	}{
		{
			name:     "Test 1",
			cases:    [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}},
			expected: "A",
		},
		{
			name:     "Test 2",
			cases:    [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
			expected: "B",
		},
		{
			name:     "Test 3",
			cases:    [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}},
			expected: "Draw",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1275.Tictactoe(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
