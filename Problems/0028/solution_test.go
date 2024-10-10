package problem0028_test

import (
	"testing"

	problem0028 "github.com/realtemirov/leetcode/Problems/0028"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []string{"sadbutsad", "sad"},
			expected: 0,
		},
		{
			name:     "Test 2",
			cases:    []string{"leetcode", "leeto"},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0028.StrStr(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
