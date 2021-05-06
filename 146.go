package leetcode

type node struct {
	next, prv *node
	key, val  int

	frequency int
}

type LRUCache struct {
	head  *node
	hash  map[int]*node
	cap   int
	dumpy *node
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	dumpy := &node{}
	head.next = dumpy
	head.prv = dumpy
	dumpy.prv = head
	dumpy.next = head
	return LRUCache{
		dumpy: dumpy,
		head:  head,
		hash:  map[int]*node{},
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	target, ok := this.hash[key]
	if !ok {
		return -1
	}

	// update priority

	target.prv.next = target.next
	target.next.prv = target.prv

	target.next = this.head.next
	this.head.next.prv = target

	this.head.next = target
	target.prv = this.head

	return target.val
}

func (this *LRUCache) Put(key int, value int) {
	target, ok := this.hash[key]
	if !ok {
		target = &node{key: key, val: value}
		this.hash[key] = target
	} else {
		target.val = value
		target.prv.next = target.next
		target.next.prv = target.prv

	}

	target.next = this.head.next
	this.head.next.prv = target
	target.prv = this.head
	this.head.next = target

	if !ok && this.cap < len(this.hash) {
		tail := this.head.prv.prv
		delete(this.hash, tail.key)
		tail.prv.next = tail.next
		tail.next.prv = tail.prv
	}

}
