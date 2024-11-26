package problem0035_test

import (
	"testing"

	problem0035 "github.com/realtemirov/leetcode/Problems/0035"
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
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			expected: 2,
		},
		{
			name: "Test 2",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			expected: 1,
		},
		{
			name: "Test 3",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0035.SearchInsert(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
