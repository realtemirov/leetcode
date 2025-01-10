package problem0916_test

import (
	"testing"

	problem0916 "github.com/realtemirov/leetcode/Problems/0916"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    [][]string
		expected []string
	}{
		{
			name: "Test 1",
			cases: [][]string{
				{"amazon", "apple", "facebook", "google", "leetcode"},
				{"e", "o"},
			},
			expected: []string{"facebook", "google", "leetcode"},
		},
		{
			name: "Test 2",
			cases: [][]string{
				{"amazon", "apple", "facebook", "google", "leetcode"},
				{"l", "e"},
			},
			expected: []string{"apple", "google", "leetcode"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0916.WordSubsets(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
