package problem0167_test

import (
	"testing"

	problem0167 "github.com/realtemirov/leetcode/Problems/0167"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums   []int
			target int
		}
		expected []int
	}{
		{
			name: "Test 1",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{1, 2},
		},
		{
			name: "Test 2",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{2, 3, 4},
				target: 6,
			},
			expected: []int{1, 3},
		},
		{
			name: "Test 3",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{-1, 0},
				target: -1,
			},
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result1 := problem0167.TwoSum(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result1, "expected: %v, result: %v", tc.expected, result1)

			result2 := problem0167.TwoSumBinarySearch(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result2, "expected: %v, result: %v", tc.expected, result2)
		})
	}
}
