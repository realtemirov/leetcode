package problem0380_test

import (
	"testing"

	problem0380 "github.com/realtemirov/leetcode/Problems/0380"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	val := problem0380.Constructor()
	require.True(t, (&val).Insert(1), "insert")
	require.False(t, (&val).Remove(2), "remove")
	require.True(t, (&val).Insert(2), "insert")
	require.True(t, func() bool {
		number := (&val).GetRandom()
		return number == 1 || number == 2
	}(), "getRandom")
	require.True(t, (&val).Remove(1), "remove")
	require.False(t, (&val).Insert(2), "insert")
	require.Equal(t, 2, (&val).GetRandom(), "getRandom")
}
