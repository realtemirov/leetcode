package problem0705

type Node struct {
	Value int
	Next  *Node
}

type MyHashSet struct {
	buckets []*Node
	count   int
}

func Constructor() MyHashSet {
	buckets := make([]*Node, 4)
	return MyHashSet{
		buckets: buckets,
		count:   0,
	}
}

func (h *MyHashSet) index(key int) int {
	return key % len(h.buckets)
}

func (h *MyHashSet) Add(key int) {
	idx := h.index(key)

	curr := h.buckets[idx]

	if curr == nil {
		h.buckets[idx] = &Node{Value: key}
		h.count++
		return
	}

	for curr.Next != nil {
		if curr.Value == key {
			return
		}

		curr = curr.Next
	}

	if curr.Value != key {
		curr.Next = &Node{Value: key}
		h.count++
	}

	if float64(h.count)/float64(len(h.buckets)) > 0.75 {
		h.resize()
	}
}

func (h *MyHashSet) resize() {
	oldBuckets := h.buckets
	h.buckets = make([]*Node, len(h.buckets)*2)
	h.count = 0

	for _, key := range oldBuckets {
		for key != nil {
			h.Add(key.Value)
		}
	}
}

func (h *MyHashSet) Remove(key int) {
	idx := h.index(key)
	curr := h.buckets[idx]

	if curr != nil && curr.Value == key {
		h.buckets[idx] = h.buckets[idx].Next
		h.count--
	}

	for curr != nil && curr.Next != nil {
		if curr.Next.Value == key {
			curr.Next = curr.Next.Next
			h.count--
		}

		curr = curr.Next
	}
}

func (h *MyHashSet) Contains(key int) bool {
	idx := h.index(key)
	curr := h.buckets[idx]

	for curr != nil {
		if curr.Value == key {
			return true
		}

		curr = curr.Next
	}

	return false
}

/**
* Your MyHashSet object will be instantiated and called as such:
* obj := Constructor();
* obj.Add(key);
* obj.Remove(key);
* param_3 := obj.Contains(key);
 */
