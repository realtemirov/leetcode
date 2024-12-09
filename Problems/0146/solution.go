package problem0146

type LRUCache struct {
	Head, Tail *Node
	M          map[int]*Node
	Capacity   int
}

type Node struct {
	Prev, Next *Node
	Key, Value int
}

func Constructor(capacity int) LRUCache {
	head := &Node{Key: 0, Value: 0}
	tail := &Node{Key: 0, Value: 0}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Head:     head,
		Tail:     tail,
		M:        make(map[int]*Node),
		Capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.M[key]
	if ok {
		cache.remove(node)
		cache.insert(node)
		return node.Value
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if _, ok := cache.M[key]; ok {
		cache.remove(cache.M[key])
	}

	if len(cache.M) == cache.Capacity {
		cache.remove(cache.Tail.Prev)
	}

	cache.insert(&Node{Key: key, Value: value})
}

func (cache *LRUCache) remove(node *Node) {
	delete(cache.M, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (cache *LRUCache) insert(node *Node) {
	cache.M[node.Key] = node
	next := cache.Head.Next
	cache.Head.Next = node
	node.Prev = cache.Head
	node.Next = next
	next.Prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
