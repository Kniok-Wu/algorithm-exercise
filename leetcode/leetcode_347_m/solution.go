package main

import (
	"fmt"
)

type MinHeap struct {
	heap   []int
	maxCap int
	record map[int]int
}

func (h *MinHeap) Insert(val int) {
	i := len(h.heap)
	h.heap = append(h.heap, val)

	for i > 0 {
		var parent int
		if i%2 == 0 {
			parent = (i - 2) / 2
		} else {
			parent = (i - 1) / 2
		}

		if h.record[h.heap[parent]] > h.record[h.heap[i]] {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		} else {
			break
		}

		i = parent
	}

	if len(h.heap) > h.maxCap {
		h.Pop()
	}
}

func (h *MinHeap) Pop() int {
	res := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	for i := 0; i < len(h.heap); {
		lchild := i*2 + 1
		rchild := i*2 + 2
		swap := -1

		if lchild < len(h.heap) {
			if rchild >= len(h.heap) { // 如果仅存在左孩子
				swap = lchild
			} else { // 如果存在右孩子
				if h.record[h.heap[lchild]] <= h.record[h.heap[rchild]] { // 左孩子较小
					swap = lchild
				} else { // 右孩子较小
					swap = rchild
				}
			}

			if h.record[h.heap[swap]] < h.record[h.heap[i]] { // 如果目标交换节点 相较于根节点更小
				h.heap[swap], h.heap[i] = h.heap[i], h.heap[swap]
			} else {
				swap = -1
			}
		}

		if swap == -1 { // 没有发生交换行为 则说明构成小根堆
			break
		}

		i = swap
	}

	return res
}

func topKFrequent(nums []int, k int) []int {
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		record[nums[i]] += 1
	}

	mh := &MinHeap{
		heap:   make([]int, 0),
		maxCap: k,
		record: record,
	}

	for key := range record {
		mh.Insert(key)
	}

	return mh.heap
}

func main() {
	fmt.Println(topKFrequent([]int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0}, 6))
}
