package problem0027_test

import (
	"testing"

	problem0027 "github.com/realtemirov/leetcode/Problems/0027"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums []int
			val  int
		}
		expected int
	}{
		{
			name: "Test 1",
			cases: struct {
				nums []int
				val  int
			}{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			expected: 2,
		},
		{
			name: "Test 2",
			cases: struct {
				nums []int
				val  int
			}{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0027.RemoveElement(tc.cases.nums, tc.cases.val)

			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
