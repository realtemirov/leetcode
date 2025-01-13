package problem1295_test

import (
	"testing"

	problem1295 "github.com/realtemirov/leetcode/Problems/1295"
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
			cases:    []int{12, 345, 2, 6, 7896},
			expected: 2,
		},
		{
			name:     "Test 2",
			cases:    []int{555, 901, 482, 1771},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem1295.FindNumbers(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
