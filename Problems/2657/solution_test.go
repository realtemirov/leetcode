package problem2657_test

import (
	"testing"

	problem2657 "github.com/realtemirov/leetcode/Problems/2657"
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
			cases:    [][]int{{1, 3, 2, 4}, {3, 1, 2, 4}},
			expected: []int{0, 2, 3, 4},
		},
		{
			name:     "Test 2",
			cases:    [][]int{{2, 3, 1}, {3, 1, 2}},
			expected: []int{0, 1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2657.FindThePrefixCommonArray1(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})

		t.Run(tc.name, func(t *testing.T) {
			result := problem2657.FindThePrefixCommonArray2(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
