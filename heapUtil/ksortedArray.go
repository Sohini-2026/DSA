package heapUtil

/*
Given an array of n elements, where each element is at most k away from its target position, devise an algorithm that sorts in O(n log k) time. For example, let us consider k is 2, an element at index 7 in the sorted array, can be at indexes 5, 6, 7, 8, 9 in the given array.

Example:
Input : arr[] = {6, 5, 3, 2, 8, 10, 9}
k = 3
Output : arr[] = {2, 3, 5, 6, 8, 9, 10} .
*/
import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func KSortedArray(arr []int, k int) ([]int, error) {
	if k > len(arr) {
		return nil, errors.New("k is larger than array size")
	}

	h := &baseDS.IntMinHeap{}
	heap.Init(h)

	result := []int{}

	// Add first k+1 elements to the heap
	for i := 0; i <= k && i < len(arr); i++ {
		heap.Push(h, arr[i])
	}

	fmt.Printf("Initial min-heap: %+v\n", *h) //[2 3 5 6]

	// Process the remaining elements
	for i := k + 1; i < len(arr); i++ {
		minElem := heap.Pop(h).(int)
		result = append(result, minElem)
		heap.Push(h, arr[i])
		fmt.Printf("Updated min-heap: %+v\n", *h)
	}

	// Extract remaining elements from the heap
	for h.Len() > 0 {
		minElem := heap.Pop(h).(int)
		result = append(result, minElem)
	}

	return result, nil
}
