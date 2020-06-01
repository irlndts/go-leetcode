package leetcode

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k    int
	heap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)

	sort.Ints(nums)
	if len(nums) < k {
		for _, n := range nums {
			heap.Push(h, n)
		}
	} else {
		for i := len(nums) - 1; i > len(nums)-k-1; i-- {
			heap.Push(h, nums[i])
		}
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.heap) >= this.k && (*this.heap)[0] < val {
		(*this.heap)[0] = val
		heap.Fix(this.heap, 0)
	}

	if len(*this.heap) < this.k {
		heap.Push(this.heap, val)
	}

	return (*this.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
