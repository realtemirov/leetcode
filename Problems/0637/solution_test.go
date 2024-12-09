package problem0637_test

import (
	"testing"

	problem0637 "github.com/realtemirov/leetcode/Problems/0637"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []any
		expected []float64
	}{
		{
			name:     "Test 1",
			cases:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: []float64{3.00000, 14.50000, 11.00000},
		},
		{
			name:     "Test 2",
			cases:    []any{3, 9, 20, 15, 7},
			expected: []float64{3.00000, 14.50000, 11.00000},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0637.AverageOfLevels(helper.SliceToTree(tc.cases))
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
