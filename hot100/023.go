package hot100

import "container/heap"

type listNodeHeap struct {
	nodes []*ListNode
}

func (h *listNodeHeap) Len() int {
	return len(h.nodes)
}

func (h *listNodeHeap) Less(i, j int) bool {
	return h.nodes[i].Val < h.nodes[j].Val
}

func (h *listNodeHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] =
		h.nodes[j], h.nodes[i]
}

// add x as element Len()
func (h *listNodeHeap) Push(x any) {
	h.nodes = append(h.nodes, x.(*ListNode))
}

// remove and return element Len() - 1.
func (h *listNodeHeap) Pop() any {
	rv := h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[0 : len(h.nodes)-1]
	return rv
}

func mergeKLists(lists []*ListNode) *ListNode {
	listHeap := &listNodeHeap{}
	for _, list := range lists {
		if list != nil {
			listHeap.nodes = append(listHeap.nodes, list)
		}
	}
	heap.Init(listHeap)
	var rv *ListNode
	var tail *ListNode
	for len(listHeap.nodes) != 0 {
		x := heap.Pop(listHeap)
		if rv == nil {
			rv = x.(*ListNode)
			tail = rv
		} else {
			tail.Next = x.(*ListNode)
			tail = tail.Next
		}
		next := x.(*ListNode).Next
		x.(*ListNode).Next = nil
		if next != nil {
			heap.Push(listHeap, next)
		}
	}
	return rv
}
