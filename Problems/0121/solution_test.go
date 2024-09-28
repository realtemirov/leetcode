package problem0121_test

import (
	"testing"

	problem0121 "github.com/realtemirov/leetcode/Problems/0121"
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
			expected: 5,
		},
		{
			name:     "Test 2",
			cases:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0121.MaxProfit(tc.cases)
			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
