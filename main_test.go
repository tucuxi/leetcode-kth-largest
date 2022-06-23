package main

import "testing"

func TestFindKthLargest(t *testing.T) {
	nums := []int{5, 3, 10, 4, 8, 7, 1, 9, 6, 2}
	for i := range nums {
		k := i + 1
		expected := 10 - i
		if actual := findKthLargest(nums, k); actual != expected {
			t.Errorf("findKthLargest(nums, %d) expected %d, got %d", k, expected, actual)
		}
	}
}
