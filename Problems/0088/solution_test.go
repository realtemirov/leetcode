package problem0088_test

import (
	"testing"

	problem0001 "github.com/realtemirov/leetcode/Problems/0088"
	"github.com/stretchr/testify/require"
)

func Test_Merge(t *testing.T) {
	testCases := []struct {
		name  string
		cases struct {
			nums1 []int
			m     int
			nums2 []int
			n     int
		}
		expected []int
	}{
		{
			name: "Test 1",
			cases: struct {
				nums1 []int
				m     int
				nums2 []int
				n     int
			}{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Test 2",
			cases: struct {
				nums1 []int
				m     int
				nums2 []int
				n     int
			}{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			expected: []int{1},
		},
		{
			name: "",
			cases: struct {
				nums1 []int
				m     int
				nums2 []int
				n     int
			}{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			problem0001.Merge(tc.cases.nums1, tc.cases.m, tc.cases.nums2, tc.cases.n)
			result := tc.cases.nums1

			require.Equal(
				t, tc.expected, result,
				"expected: %v, result: %v", tc.expected, result,
			)
		})
	}
}
