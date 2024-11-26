package problem0034_test

import (
	"testing"

	problem0034 "github.com/realtemirov/leetcode/Problems/0034"
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
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			expected: []int{3, 4},
		},
		{
			name: "Test 2",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			expected: []int{-1, -1},
		},
		{
			name: "Test 3",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{},
				target: 0,
			},
			expected: []int{-1, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0034.SearchRange(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}

func TestLowerBound(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums   []int
			target int
		}
		expected int
	}{
		{
			name: "Target found",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 2, 3, 4, 5},
				target: 2,
			},
			expected: 1,
		},
		{
			name: "Target not found - insert position",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 4, 5, 6},
				target: 3,
			},
			expected: 2,
		},
		{
			name: "Target greater than all elements",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 3},
				target: 10,
			},
			expected: 3,
		},
		{
			name: "Empty array",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{},
				target: 1,
			},
			expected: 0,
		},
		{
			name: "Target smaller than all elements",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{5, 6, 7},
				target: 2,
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0034.LowerBound(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result, "expected: %v, got: %v", tc.expected, result)
		})
	}
}

func TestUpperBound(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums   []int
			target int
		}
		expected int
	}{
		{
			name: "Target found",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 2, 3, 4, 5},
				target: 2,
			},
			expected: 3,
		},
		{
			name: "Target not found - insert position",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 4, 5, 6},
				target: 3,
			},
			expected: 2,
		},
		{
			name: "Target greater than all elements",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 3},
				target: 10,
			},
			expected: 3,
		},
		{
			name: "Empty array",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{},
				target: 1,
			},
			expected: 0,
		},
		{
			name: "Target smaller than all elements",
			cases: struct {
				nums   []int
				target int
			}{
				nums:   []int{5, 6, 7},
				target: 2,
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0034.UpperBound(tc.cases.nums, tc.cases.target)
			require.Equal(t, tc.expected, result, "expected: %v, got: %v", tc.expected, result)
		})
	}
}
