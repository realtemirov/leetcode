package problem0724_test

import (
	"testing"

	problem0724 "github.com/realtemirov/leetcode/Problems/0724"
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
			cases:    []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "Test 3",
			cases:    []int{2, 1, -1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0724.PivotIndex(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
