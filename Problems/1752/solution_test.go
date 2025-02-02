package problem1752_test

import (
	"testing"

	problem1752 "github.com/realtemirov/leetcode/Problems/1752"
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
			cases:    []int{3, 4, 5, 1, 2},
			expected: true,
		},
		{
			name:     "Test 2",
			cases:    []int{2, 1, 3, 4},
			expected: false,
		},
		{
			name:     "Test 3",
			cases:    []int{1, 2, 3},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1752.Check1(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})

		t.Run(tc.name, func(t *testing.T) {
			result := problem1752.Check2(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
