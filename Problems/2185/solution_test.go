package problem2185_test

import (
	"testing"

	problem2185 "github.com/realtemirov/leetcode/Problems/2185"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			words []string
			pref  string
		}
		expected int
	}{
		{
			name: "Test 1",
			cases: struct {
				words []string
				pref  string
			}{
				words: []string{"pay", "attention", "practice", "attend"},
				pref:  "at",
			},
			expected: 2,
		},
		{
			name: "Test 2",
			cases: struct {
				words []string
				pref  string
			}{
				words: []string{"leetcode", "win", "loops", "success"},
				pref:  "code",
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem2185.PrefixCount(tc.cases.words, tc.cases.pref)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
