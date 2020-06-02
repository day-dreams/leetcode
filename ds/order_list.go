package ds

import "fmt"

type listNode struct {
	val  int
	next *listNode
}
type OrderList struct {
	head *listNode
	cnt  int
}

func NewOrderList() *OrderList {
	return &OrderList{}
}

func (ol *OrderList) Add(val int) {

	defer func() { ol.cnt++ }()

	if ol.head == nil {
		ol.head = &listNode{
			val:  val,
			next: nil,
		}
		return
	}

	var (
		prv *listNode
		cur = ol.head
	)
	for cur != nil && cur.val < val {
		prv = cur
		cur = cur.next
	}

	if cur == ol.head {
		ol.head = &listNode{val: val, next: ol.head}
	} else if cur == nil {
		prv.next = &listNode{val: val, next: nil}
	} else {
		prv.next = &listNode{val: val, next: cur}
	}

}

func (ol *OrderList) Traverse(lambda func(val int)) {
	cur := ol.head
	for cur != nil {
		lambda(cur.val)
		cur = cur.next
	}
}

func (ol *OrderList) KMax(k int) (val int, err error) {
	if ol.cnt < k {
		return 0, fmt.Errorf("want %dth max elemet, but only %d", k, ol.cnt)
	}

	cur := ol.head
	index := ol.cnt - k
	for index > 0 {
		index--
		cur = cur.next
	}

	return cur.val, nil
}
