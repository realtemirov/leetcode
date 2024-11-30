package problem0075_test

import (
	"testing"

	problem0075 "github.com/realtemirov/leetcode/Problems/0075"
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
			cases:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Test 2",
			cases:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			problem0075.SortColors(tc.cases)
			require.Equal(t, tc.expected, tc.cases, "expected: %v, result: %v", tc.expected, tc.cases)
		})
	}
}
