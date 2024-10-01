package problem0045_test

import (
	"testing"

	problem0045 "github.com/realtemirov/leetcode/Problems/0045"
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
			cases:    []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "Test 2",
			cases:    []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "Test 3",
			cases:    []int{0},
			expected: 0,
		},
		{
			name:     "Test 4",
			cases:    []int{2, 1},
			expected: 1,
		},
		{
			name:     "Test 5",
			cases:    []int{1, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "Test 6",
			cases:    []int{5, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Test 7",
			cases:    []int{3, 2, 1, 0, 4},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0045.Jump(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
