package problem0006_test

import (
	"testing"

	problem0006 "github.com/realtemirov/leetcode/Problems/0006"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			word string
			rows int
		}
		expected string
	}{
		{
			name: "Test 1",
			cases: struct {
				word string
				rows int
			}{
				word: "PAYPALISHIRING",
				rows: 3,
			},
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name: "Test 2",
			cases: struct {
				word string
				rows int
			}{
				word: "PAYPALISHIRING",
				rows: 4,
			},
			expected: "PINALSIGYAHRPI",
		},
		{
			name: "Test 3",
			cases: struct {
				word string
				rows int
			}{
				word: "A",
				rows: 1,
			},
			expected: "A",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0006.Convert(tc.cases.word, tc.cases.rows)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
