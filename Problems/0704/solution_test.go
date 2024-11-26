package problem0704_test

import (
	"testing"

	problem0704 "github.com/realtemirov/leetcode/Problems/0704"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums   []int
			target int
		}
		expected int
	}{
		{
			name: "Test 1",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			expected: 4,
		},
		{
			name: "Test 2",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0704.Search(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
