package hot100

import (
	"fmt"
	"testing"
)

func print(c *LRUCache) {
	s := ""
	for head := c.head; head != nil; head = head.next {
		s += fmt.Sprintf("%v,%v ", head.key, head.value)
	}
	fmt.Printf("%s\n", s)
}

func TestConstructor(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	print(&cache)
	cache.Put(2, 2)
	print(&cache)
	cache.Get(1)
	print(&cache)
	cache.Put(3, 3)
	print(&cache)
	cache.Get(2)
	print(&cache)
	cache.Put(4, 4)
	print(&cache)
	cache.Get(1)
	print(&cache)
	cache.Get(3)
	print(&cache)
	cache.Get(4)
	print(&cache)
}
