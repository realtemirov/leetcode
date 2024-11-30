package problem0001_test

import (
	"testing"

	problem0001 "github.com/realtemirov/leetcode/Problems/0001"
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
			expected: []int{0, 1},
		},
		{
			name: "Test 2",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			name: "Test 3",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0001.TwoSum(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
