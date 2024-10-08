package problem0011_test

import (
	"testing"

	problem0011 "github.com/realtemirov/leetcode/Problems/0011"
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
			cases:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0011.MaxArea(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
