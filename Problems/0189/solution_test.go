package problem0189_test

import (
	"testing"

	problem0189 "github.com/realtemirov/leetcode/Problems/0189"
	"github.com/stretchr/testify/require"
)

func Test_Rotate(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums []int
			val  int
		}
		expected []int
	}{
		{
			name: "Test 1",
			cases: struct {
				nums []int
				val  int
			}{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				val:  3,
			},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Test 2",
			cases: struct {
				nums []int
				val  int
			}{
				nums: []int{-1, -100, 3, 99},
				val:  2,
			},
			expected: []int{3, 99, -1, -100},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			problem0189.Rotate(tc.cases.nums, tc.cases.val)
			result := tc.cases.nums
			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
