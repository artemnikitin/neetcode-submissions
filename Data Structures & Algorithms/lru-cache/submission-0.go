type LRUCache struct {
	capacity int
	cache    map[int]*Node
	left     *Node
	right    *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		left:     &Node{},
		right:    &Node{},
	}
	lruCache.left.next = lruCache.right
	lruCache.right.prev = lruCache.left
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.insert(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		this.remove(node)
		delete(this.cache, key)
	}
	node = &Node{key: key, value: value}
	this.cache[key] = node
	this.insert(node)
	if len(this.cache) > this.capacity {
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func (this *LRUCache) insert(node *Node) {
	prev := this.right.prev
	next := this.right
	prev.next = node
	next.prev = node
	node.prev = prev
	node.next = next
}

func (this *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}
