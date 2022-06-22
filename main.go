package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := MaxHeap(nums)
	heap.Init(&h)
	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

func main() {
	nums := []int{5, 3, 4, 8, 7, 1, 9, 6, 2}
	a := findKthLargest(nums, 3)
	fmt.Printf("nums=%v, third largest element is %d\n", nums, a)
	b := findKthLargest(nums, 8)
	fmt.Printf("nums=%v, eighth largest element is %d\n", nums, b)
}
