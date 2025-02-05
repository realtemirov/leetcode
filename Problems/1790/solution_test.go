package problem1790_test

import (
	"testing"

	problem1790 "github.com/realtemirov/leetcode/Problems/1790"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected bool
	}{
		{
			name:     "Test 1",
			cases:    []string{"bank", "kanb"},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []string{"attack", "defend"},
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    []string{"kelb", "kelb"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1790.AreAlmostEqual(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
