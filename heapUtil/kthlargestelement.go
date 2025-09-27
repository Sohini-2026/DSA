package heapUtil

/*Given an array and a number k where k is smaller than size of array, we need to find the k’th smallest element in the given array. It is given that all array elements are distinct.
Example:
Input: arr[] = {7, 10, 4, 3, 20, 15}
k = 2
Output: 10 .*/

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func KthLargest(arr []int, k int) (int, error) {
	if k > len(arr) {
		return -1, errors.New("k is larger than array size")
	}

	h := &baseDS.IntMinHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}

	fmt.Printf("Initial min-heap: %+v\n", *h) //[10 7 4]

	//([]int{7, 10, 4, 3, 20, 15}, 3)

	for i := k; i < len(arr); i++ {
		if arr[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
			fmt.Printf("Updated min-heap: %+v\n", *h)
		}
	}

	return (*h)[0], nil
}
