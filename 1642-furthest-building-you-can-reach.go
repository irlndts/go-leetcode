package leetcode

import "container/heap"

// https://leetcode.com/problems/furthest-building-you-can-reach/

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}



func furthestBuilding(heights []int, bricks, ladders int) int {
	h := &MaxHeap{}
	heap.Init(h)

	size := len(heights)
	if size == 1 {
		return 0
	}

	for i := 1; i < size; i++ {
		diff := heights[i] - heights[i-1]

		if diff > 0 {
			if bricks >= diff {
				bricks -= diff
				heap.Push(h, diff)
				continue
			}

			if ladders == 0 {
				return i - 1
			}

			ladders--
			heap.Push(h, diff)
			pop := heap.Pop(h).(int)

			if pop != diff {
				bricks += pop - diff
			}
		}
	}

	return size - 1
}