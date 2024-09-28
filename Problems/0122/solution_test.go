package problem0122_test

import (
	"testing"

    problem0122 "github.com/realtemirov/leetcode/Problems/0122"
	"github.com/stretchr/testify/require"
)

func Test_MaxProfit(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "Test 2",
			cases:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Test 3",
			cases:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0122.MaxProfit(tc.cases)
			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
