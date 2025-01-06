package problem0067_test

import (
	"testing"

	problem0067 "github.com/realtemirov/leetcode/Problems/0067"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []string
		expected string
	}{
		{
			name:     "Test 1",
			cases:    []string{"11", "1"},
			expected: "100",
		},
		{
			name:     "Test 2",
			cases:    []string{"1010", "1011"},
			expected: "10101",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0067.AddBinary(tc.cases[0], tc.cases[1])
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
