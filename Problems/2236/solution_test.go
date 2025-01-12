package problem2236_test

import (
	"testing"

	problem2236 "github.com/realtemirov/leetcode/Problems/2236"
	"github.com/realtemirov/leetcode/helper"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []any
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    []any{10, 4, 6},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []any{5, 3, 1},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2236.CheckTree(helper.SliceToTree(tc.cases))
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
