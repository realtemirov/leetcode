package problem0682_test

import (
	"testing"

	problem0682 "github.com/realtemirov/leetcode/Problems/0682"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []string{"5", "2", "C", "D", "+"},
			expected: 30,
		},
		{
			name:     "Test 2",
			cases:    []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			expected: 27,
		},
		{
			name:     "Test 3",
			cases:    []string{"1", "C"},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0682.CalPoints(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
