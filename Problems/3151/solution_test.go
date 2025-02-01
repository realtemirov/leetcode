package problem3151_test

import (
	"testing"

	problem3151 "github.com/realtemirov/leetcode/Problems/3151"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    []int{1},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{2, 1, 4},
			expected: true,
		},
		{
			name:     "Test 3",
			cases:    []int{4, 3, 1, 6},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem3151.IsArraySpecial(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
