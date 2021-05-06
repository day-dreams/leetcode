package leetcode

type LFUCache struct {
	cap  int
	hash map[int]*node
	head *node
}

func LFUConstructor(capacity int) LFUCache {
	dump := &node{}
	head := &node{}
	head.next = dump
	head.prv = dump
	dump.next = head
	dump.prv = head
	return LFUCache{
		cap:  capacity,
		hash: map[int]*node{},
		head: head,
	}
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}

	target, ok := this.hash[key]
	if !ok {
		return -1
	}

	target.frequency++
	cur := target.prv
	for target.frequency >= cur.frequency && cur != this.head {
		cur = cur.prv
	}

	if cur.next == target {
		return target.val
	}

	target.next.prv = target.prv
	target.prv.next = target.next
	target.next = nil
	target.prv = nil

	cur.next.prv = target
	target.next = cur.next
	target.prv = cur
	cur.next = target

	return target.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	target, ok := this.hash[key]
	if !ok && len(this.hash) == this.cap {
		lfu := this.head.prv.prv
		delete(this.hash, lfu.key)
		lfu.prv.next = lfu.next
		lfu.next.prv = lfu.prv
	}
	if !ok {
		target = &node{key: key, val: value, frequency: 0}
	}
	if ok {
		target.val = value
	}
	target.frequency++
	this.hash[key] = target

	cur := this.head.prv.prv
	for cur.frequency <= target.frequency && cur != this.head {
		cur = cur.prv
	}

	if cur.next == target {
		return
	}

	if ok {
		target.prv.next = target.next
		target.next.prv = target.prv
	}

	target.next = cur.next
	cur.next.prv = target
	cur.next = target
	target.prv = cur
}
