package problem0068_test

import (
	"testing"

	problem0068 "github.com/realtemirov/leetcode/Problems/0068"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			words    []string
			maxWidth int
		}
		expected []string
	}{
		{
			name: "Test 1",
			cases: struct {
				words    []string
				maxWidth int
			}{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			name: "Test 2",
			cases: struct {
				words    []string
				maxWidth int
			}{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			expected: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			name: "Test 3",
			cases: struct {
				words    []string
				maxWidth int
			}{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20,
			},
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0068.FullJustify(tc.cases.words, tc.cases.maxWidth)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
