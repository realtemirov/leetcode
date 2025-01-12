package problem0412_test

import (
	"testing"

	problem0412 "github.com/realtemirov/leetcode/Problems/0412"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		cases    int
		expected []string
	}{
		{
			name:     "Test 1",
			cases:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "Test 2",
			cases:    5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:     "Test 3",
			cases:    15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0412.FizzBuzz(tc.cases)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
