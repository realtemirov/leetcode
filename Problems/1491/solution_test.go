package problem1491_test

import (
	"testing"

	problem1491 "github.com/realtemirov/leetcode/Problems/1491"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected float64
	}{
		{
			name:     "Test 1",
			cases:    []int{4000, 3000, 1000, 2000},
			expected: 2500.00000,
		},
		{
			name:     "Test 2",
			cases:    []int{1000, 2000, 3000},
			expected: 2000.00000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1491.Average(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
