package datstr

type lruNode[V any] struct {
	val  V
	next *lruNode[V]
	prev *lruNode[V]
}

func createNode[V any](value V) *lruNode[V] {
	return &lruNode[V]{
		val:  value,
		next: nil, prev: nil,
	}
}

type LRU[K comparable, V any] struct {
	lenght        int
	cap           int
	head          *lruNode[V]
	tail          *lruNode[V]
	lookup        map[K]*lruNode[V]
	reverseLookup map[*lruNode[V]]K
}

func NewLRU[K comparable, V any](cap int) *LRU[K, V] {
	return &LRU[K, V]{
		lenght:        0,
		cap:           cap,
		head:          nil,
		tail:          nil,
		lookup:        make(map[K]*lruNode[V], cap),
		reverseLookup: make(map[*lruNode[V]]K),
	}
}

func (l *LRU[K, V]) Update(key K, value V) any {
	// does it exist?
	node, found := l.lookup[key]
	if !found {
		node = createNode(value)
		l.lenght++
		l.prepend(node)
		l.trimCache()

		l.lookup[key] = node
		l.reverseLookup[node] = key
	} else {
		l.detach(node)
		l.prepend(node)
		node.val = value
	}
	// if it doesn't we need to insert
	// - check capacity and evict if over
	// if it does, we need to update to the front of the list
	// and update the value
	return node.val
}

func (l *LRU[K, V]) Get(key K) any {
	// check the cache for existence
	node, found := l.lookup[key]
	if !found {
		return nil
	}

	// update the value we found and move it to the front
	l.detach(node)
	l.prepend(node)

	// return the out the value found or nil if not found
	return node.val
}

func (l *LRU[K, V]) detach(node *lruNode[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}
	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *LRU[K, V]) prepend(node *lruNode[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node

	l.head = node
}

func (l *LRU[K, V]) trimCache() {
	if l.lenght <= l.cap {
		return
	}

	tail := l.tail
	l.detach(tail)

	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.lenght--
}
