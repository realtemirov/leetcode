package problem2235_test

import (
	"testing"

	problem2235 "github.com/realtemirov/leetcode/Problems/2235"
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
			cases:    []int{12, 5},
			expected: 17,
		},
		{
			name:     "Test 2",
			cases:    []int{-10, 4},
			expected: -6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2235.Sum(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
