package problem0146_test

import (
	"testing"

	problem0146 "github.com/realtemirov/leetcode/Problems/0146"
	"github.com/stretchr/testify/require"
)

func TestLRUCacheSequence(t *testing.T) {
	// Initialize the LRUCache with a capacity of 2
	cache := problem0146.Constructor(2)

	// Test Case 1: Put (1,1)
	cache.Put(1, 1) // cache = [(1, 1)]

	// Test Case 2: Put (2,2)
	cache.Put(2, 2) // cache = [(2, 2), (1, 1)]

	// Test Case 3: Get (1)
	require.Equal(t, 1, cache.Get(1), "Expected value for key 1 is 1")
	// cache = [(1, 1), (2, 2)]

	// Test Case 4: Put (3,3)
	cache.Put(3, 3) // cache = [(3, 3), (1, 1)], evicts (2, 2)

	// Test Case 5: Get (2)
	require.Equal(t, -1, cache.Get(2), "Expected value for key 2 is -1 (evicted)")

	// Test Case 6: Put (4,4)
	cache.Put(4, 4) // cache = [(4, 4), (3, 3)], evicts (1, 1)

	// Test Case 7: Get (1)
	require.Equal(t, -1, cache.Get(1), "Expected value for key 1 is -1 (evicted)")

	// Test Case 8: Get (3)
	require.Equal(t, 3, cache.Get(3), "Expected value for key 3 is 3")
	// cache = [(3, 3), (4, 4)]

	// Test Case 9: Get (4)
	require.Equal(t, 4, cache.Get(4), "Expected value for key 4 is 4")
	// cache = [(4, 4), (3, 3)]
}
