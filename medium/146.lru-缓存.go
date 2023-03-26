/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package main

// @lc code=start
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := &DLinkedNode{
			key:   key,
			value: value,
		}
		this.cache[key] = node
		this.addToHead(node)
		this.capacity--
		if this.capacity < 0 {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.capacity++
		}
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
