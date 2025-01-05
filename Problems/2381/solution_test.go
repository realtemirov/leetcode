package problem2381_test

import (
	"testing"

	problem2381 "github.com/realtemirov/leetcode/Problems/2381"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			s      string
			shifts [][]int
		}
		expected string
	}{
		{
			name: "Test 1",
			cases: struct {
				s      string
				shifts [][]int
			}{
				s:      "abc",
				shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}},
			},
			expected: "ace",
		},
		{
			name: "Test 2",
			cases: struct {
				s      string
				shifts [][]int
			}{
				s:      "dztz",
				shifts: [][]int{{0, 0, 0}, {1, 1, 1}},
			},
			expected: "catz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2381.ShiftingLetters(tc.cases.s, tc.cases.shifts)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
