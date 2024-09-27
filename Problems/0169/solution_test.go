package problem0169_test

import (
	"testing"

	problem0169 "github.com/realtemirov/leetcode/Problems/0169"
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
			cases:    []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "Test 2",
			cases:    []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0169.MajorityElement(tc.cases)

			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
