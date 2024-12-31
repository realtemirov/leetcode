package problem0860_test

import (
	"testing"

	problem0860 "github.com/realtemirov/leetcode/Problems/0860"
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
			cases:    []int{5,5,5,10,20},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{5,5,10,10,20},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0860.LemonadeChange(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
