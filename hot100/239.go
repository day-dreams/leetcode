package hot100

import (
	"container/heap"
	"sort"
)

type windowValue struct {
	value int
	index int
}

var numBackup []int

type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool {
	return numBackup[h.IntSlice[i]] > numBackup[h.IntSlice[j]]
}

func (h *hp) Push(value interface{}) {
	h.IntSlice = append(h.IntSlice, value.(int))
}

func (h *hp) Pop() interface{} {
	rv := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return rv
}

// 我真不知道为啥自己的堆写错了？？？？
func maxSlidingWindow(nums []int, k int) []int {
	var (
		rv      = []int{}
		maxHeap = &hp{}
	)
	numBackup = nums
	heap.Init(maxHeap)
	for begin := 0; begin < len(nums); begin++ {
		heap.Push(maxHeap, begin)
		if begin < k-1 {
			continue
		}

		if begin == k-1 {
			rv = append(rv, numBackup[maxHeap.IntSlice[0]])
			continue
		}

		for {
			if maxHeap.IntSlice[0] <= begin-k {
				heap.Pop(maxHeap)
				continue
			}
			rv = append(rv, numBackup[maxHeap.IntSlice[0]])
			break
		}
	}
	return rv
}

// 自己手撸一个最大堆，为啥不对？？？
func maxSlidingWindow2(nums []int, k int) []int {
	var (
		rv = []int{}

		// 最大堆
		maxHeap          = make([]*windowValue, 0)
		fixMaxHeapOnPush = func() {
			// 最后一个元素是刚插入的，要把它上升到合适的位置去
			for currentIndex := len(maxHeap) - 1; ; {
				parentIndex := (currentIndex - 1) / 2
				if maxHeap[parentIndex].value < maxHeap[currentIndex].value {
					maxHeap[parentIndex], maxHeap[currentIndex] =
						maxHeap[currentIndex], maxHeap[parentIndex]
					currentIndex = parentIndex
					continue
				}
				break
			}
		}
		fixMaxHeapOnPop = func() {
			// 第一个元素是最小的，现在我们要把它下沉到合适位置
			for currentIndex := 0; currentIndex < len(maxHeap); {
				leftChild := currentIndex*2 + 1
				rightChild := currentIndex*2 + 2
				if leftChild >= len(maxHeap) {
					break
				}

				if rightChild >= len(maxHeap) {
					// 没有右孩子
					if maxHeap[currentIndex].value < maxHeap[leftChild].value {
						maxHeap[currentIndex], maxHeap[leftChild] =
							maxHeap[leftChild], maxHeap[currentIndex]
						currentIndex = leftChild
						continue
					}
					// 到位了
					break
				}

				if maxHeap[currentIndex].value >= maxHeap[leftChild].value &&
					maxHeap[currentIndex].value >= maxHeap[rightChild].value {
					break
				}

				if maxHeap[leftChild].value >= maxHeap[rightChild].value &&
					maxHeap[leftChild].value >= maxHeap[currentIndex].value {
					maxHeap[leftChild], maxHeap[currentIndex] =
						maxHeap[currentIndex], maxHeap[leftChild]
					currentIndex = leftChild
					continue
				}

				maxHeap[rightChild], maxHeap[currentIndex] =
					maxHeap[currentIndex], maxHeap[rightChild]
				currentIndex = rightChild
			}

		}
	)

	for begin := 0; begin < len(nums); begin++ {
		maxHeap = append(maxHeap, &windowValue{value: nums[begin], index: begin})
		fixMaxHeapOnPush()

		if begin < k-1 {
			continue
		}

		for {
			if maxHeap[0].index <= begin-k {
				maxHeap[0] = maxHeap[len(maxHeap)-1]
				maxHeap = maxHeap[0 : len(maxHeap)-1]
				fixMaxHeapOnPop()
				continue
			}

			// fmt.Printf("max value:%v index:%v current:%v\n", maxHeap[0].value, maxHeap[0].index, begin)
			rv = append(rv, maxHeap[0].value)
			break
		}
	}

	return rv
}

// 暴力 超时
func maxSlidingWindow1(nums []int, k int) []int {
	var (
		rv = []int{}
	)

	for i := 0; i+k <= len(nums); i++ {
		rv = append(rv, max(nums[i:i+k]))
	}
	return rv
}

func max(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
