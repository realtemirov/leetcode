package problem0050_test

import (
	"testing"

	problem0050 "github.com/realtemirov/leetcode/Problems/0050"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			x float64
			n int
		}
		expected float64
	}{
		{
			name: "Test 1",
			cases: struct {
				x float64
				n int
			}{
				x: 2.00000,
				n: 10,
			},
			expected: 1024.00000,
		},
		{
			name: "Test 2",
			cases: struct {
				x float64
				n int
			}{
				x: 2.10000,
				n: 3,
			},
			expected: 9.26100,
		},
		{
			name: "Test 3",
			cases: struct {
				x float64
				n int
			}{
				x: 2.00000,
				n: -2,
			},
			expected: 0.25000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem0050.MyPow(tc.cases.x, tc.cases.n)
			require.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
