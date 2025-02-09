package main

type CacheNode struct {
	Key  int
	Next *CacheNode
	// Prev *CacheNode
}

type LRUCache struct {
	c    int
	m    map[int]int
	head *CacheNode
	tail *CacheNode
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		return LRUCache{}
	}
	return LRUCache{
		c: capacity,
		m: make(map[int]int, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	if val, exists := c.m[key]; exists {
		return val
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	if len(c.m) == 0 {
		c.m[key] = value
		c.head = &CacheNode{Key: key}
	} else if len(c.m) == c.c {
		delete(c.m, c.head.Key)
		c.head = c.head.Next
	} else {
		c.m[key] = value
		c.head.Next = &CacheNode{Key: key}
	}
}
