package problem0135_test

import (
	"testing"

	problem0135 "github.com/realtemirov/leetcode/Problems/0135"
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
			cases:    []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 2},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0135.Candy(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
