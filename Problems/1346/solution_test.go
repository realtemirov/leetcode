package problem1346_test

import (
	"testing"

	problem1346 "github.com/realtemirov/leetcode/Problems/1346"
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
			cases:    []int{10, 2, 5, 3},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{3, 1, 7, 11},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1346.CheckIfExist(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
