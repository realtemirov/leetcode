package problem0219_test

import (
	"testing"

	problem0219 "github.com/realtemirov/leetcode/Problems/0219"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums []int
			k    int
		}
		expected bool
	}{
		{
			name: "Test 1",
			cases: struct {
				nums []int
				k    int
			}{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			expected: true,
		},
		{
			name: "Test 2",
			cases: struct {
				nums []int
				k    int
			}{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			expected: true,
		},
		{
			name: "Test 3",
			cases: struct {
				nums []int
				k    int
			}{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0219.ContainsNearbyDuplicate(tc.cases.nums, tc.cases.k)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
