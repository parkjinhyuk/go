package codingTest

type LRUCache struct {
	cache    map[int]*node
	capacity int
	head     *node
	tail     *node
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		map[int]*node{},
		capacity,
		nil,
		nil,
	}
}

func (this *LRUCache) Get(key int) int {
	n, exists := this.cache[key]
	if !exists {
		return -1
	}

	this.updateCache(key, n)
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	this.upsertCache(key, value)

	if len(this.cache) > this.capacity {
		this.deleteCache()
	}
}

func (this *LRUCache) updateCache(key int, n *node) {
	if this.head == n {
		return
	}

	if this.tail == n {
		this.tail = n.prev
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}

	n.prev = nil
	n.next = this.head
	this.head.prev = n
	this.head = n
}

func (this *LRUCache) upsertCache(key int, value int) {
	n, exists := this.cache[key]
	if exists {
		n.value = value
		this.updateCache(key, n)
		return
	}

	this.insertCache(key, value)
}

func (this *LRUCache) insertCache(key int, value int) {
	newHead := &node{
		key,
		value,
		nil,
		this.head,
	}

	if this.head == nil {
		this.head = newHead
		this.tail = newHead
	} else {
		this.head.prev = newHead
		this.head = newHead
	}

	this.cache[key] = newHead
}

func (this *LRUCache) deleteCache() {
	if this.tail == nil {
		return
	}

	delete(this.cache, this.tail.key)
	this.tail.prev.next = nil
	this.tail = this.tail.prev
}
