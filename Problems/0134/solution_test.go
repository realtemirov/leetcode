package problem0134_test

import (
	"testing"

	problem0134 "github.com/realtemirov/leetcode/Problems/0134"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			gas  []int
			cost []int
		}
		expected int
	}{
		{
			name: "Test 1",
			cases: struct {
				gas  []int
				cost []int
			}{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			expected: 3,
		},
		{
			name: "Test 2",
			cases: struct {
				gas  []int
				cost []int
			}{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			expected: -1,
		},
		{
			name: "Test 3",
			cases: struct {
				gas  []int
				cost []int
			}{
				gas:  []int{5, 8, 2, 8},
				cost: []int{6, 5, 6, 6},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0134.CanCompleteCircuit(tc.cases.gas, tc.cases.cost)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
