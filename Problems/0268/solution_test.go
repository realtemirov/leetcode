package problem0268_test

import (
	"testing"

	problem0268 "github.com/realtemirov/leetcode/Problems/0268"
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
			cases:    []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "Test 2",
			cases:    []int{0, 1},
			expected: 2,
		},
		{
			name:     "Test 3",
			cases:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0268.MissingNumber(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
