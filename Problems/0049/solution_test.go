package problem0049_test

import (
	"testing"

	problem0049 "github.com/realtemirov/leetcode/Problems/0049"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected [][]string
	}{
		{
			name:     "Test 1",
			cases:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "Test 2",
			cases:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Test 3",
			cases:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0049.GroupAnagrams(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
