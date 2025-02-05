package problem1800_test

import (
	"testing"

	problem1800 "github.com/realtemirov/leetcode/Problems/1800"
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
			cases:    []int{10, 20, 30, 5, 10, 50},
			expected: 65,
		},
		{
			name:     "Test 2",
			cases:    []int{10, 20, 30, 40, 50},
			expected: 150,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1800.MaxAscendingSum(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
