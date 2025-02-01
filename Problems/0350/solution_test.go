package problem0350_test

import (
	"testing"

	problem0350 "github.com/realtemirov/leetcode/Problems/0350"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    [][]int
		expected []int
	}{
		{
			name:     "Test 1",
			cases:    [][]int{{1, 2, 2, 1}, {2, 2}},
			expected: []int{2, 2},
		},
		{
			name:     "Test 2",
			cases:    [][]int{{4, 9, 5}, {9, 4, 9, 8, 4}},
			expected: []int{9, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0350.Intersect(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
