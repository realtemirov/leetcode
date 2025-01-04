package problem2559_test

import (
	"testing"

	problem2559 "github.com/realtemirov/leetcode/Problems/2559"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			words   []string
			queries [][]int
		}
		expected []int
	}{
		{
			name: "Test 1",
			cases: struct {
				words   []string
				queries [][]int
			}{
				words:   []string{"aba", "bcb", "ece", "aa", "e"},
				queries: [][]int{{0, 2}, {1, 4}, {1, 1}},
			},
			expected: []int{2, 3, 0},
		},
		{
			name: "Test 2",
			cases: struct {
				words   []string
				queries [][]int
			}{
				words:   []string{"a", "e", "i"},
				queries: [][]int{{0, 2}, {0, 1}, {2, 2}},
			},
			expected: []int{3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2559.VowelStrings(tc.cases.words, tc.cases.queries)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
