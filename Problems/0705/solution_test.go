package problem0705_test

import (
	"testing"

	problem0705 "github.com/realtemirov/leetcode/Problems/0705"
	"github.com/stretchr/testify/require"
)

func TestMyHashSet(t *testing.T) {
	myHashSet := problem0705.Constructor()

	// Test Case 1: Add 1
	myHashSet.Add(1) // set = [1]
	require.True(t, myHashSet.Contains(1), "Expected 1 to be in the set")

	// Test Case 2: Add 2
	myHashSet.Add(2) // set = [1, 2]
	require.True(t, myHashSet.Contains(2), "Expected 2 to be in the set")

	// Test Case 3: Check if 1 is in the set
	require.True(t, myHashSet.Contains(1), "Expected 1 to be in the set")

	// Test Case 4: Check if 3 is in the set
	require.False(t, myHashSet.Contains(3), "Expected 3 to not be in the set")

	// Test Case 5: Add 2 again
	myHashSet.Add(2) // set = [1, 2]
	require.True(t, myHashSet.Contains(2), "Expected 2 to be in the set")

	// Test Case 6: Remove 2
	myHashSet.Remove(2) // set = [1]
	require.False(t, myHashSet.Contains(2), "Expected 2 to not be in the set after removal")
}
