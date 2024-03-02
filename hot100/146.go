package hot100

type Item struct {
	key, value int
	prv, next  *Item
}

type LRUCache struct {
	key2item   map[int]*Item
	head, tail *Item
	cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		key2item: make(map[int]*Item),
		cap:      capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	item, ok := this.key2item[key]
	if !ok {
		return -1
	}
	this.moveToHead(key)
	return item.value
}

func (this *LRUCache) moveToHead(key int) {
	item := this.key2item[key]
	if item == this.head || len(this.key2item) == 1 {
		return
	}

	if item == this.tail {
		this.tail.prv.next, this.tail, this.head, item.prv, item.next, this.head.prv =
			nil, this.tail.prv, item, nil, this.head, item
		return
	}

	// head -> .... -> item -> ... -> tail
	item.next, this.head.prv, item.prv, this.head, this.head.prv, item.next.prv, item.prv.next =
		this.head, item, nil, item, item, item.prv, item.next
}

func (this *LRUCache) Put(key int, value int) {
	item, ok := this.key2item[key]
	if !ok {
		if this.head == nil {
			this.head = &Item{
				key:   key,
				value: value,
			}
			this.tail = this.head
			this.key2item[key] = this.head
			return
		}

		item := &Item{
			key:   key,
			value: value,
			next:  this.head,
		}
		this.head, this.head.prv = item, item
		this.key2item[key] = item
		if len(this.key2item) > this.cap {
			delete(this.key2item, this.tail.key)
			this.tail, this.tail.prv = this.tail.prv, nil
		}
		return
	}

	item.value = value
	// TODO(put 算使用吗)
	this.moveToHead(key)
}
