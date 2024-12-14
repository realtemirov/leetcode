package problem0066_test

import (
	"testing"

	problem0066 "github.com/realtemirov/leetcode/Problems/0066"
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
			cases:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "Test 2",
			cases:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "Test 3",
			cases:    []int{9},
			expected: []int{1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0066.PlusOne(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
