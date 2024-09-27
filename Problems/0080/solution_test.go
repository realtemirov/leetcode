package problem0080_test

import (
	"testing"

	problem0080 "github.com/realtemirov/leetcode/Problems/0080"
	"github.com/stretchr/testify/require"
)

func Test_RemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		cases    []int
		expected int
	}{
		{
			name:     "Test 1",
			cases:    []int{1, 1, 1, 2, 2, 3},
			expected: 5,
		},
		{
			name:     "Test 2",
			cases:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0080.RemoveDuplicates(tc.cases)

			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
